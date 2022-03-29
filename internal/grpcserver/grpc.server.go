package grpcserver

// Good info on gRPC and auth here: https://jbrandhorst.com/post/grpc-auth/

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quantstop/quantstopterminal/internal"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/pkg/system/crypto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
	"net/http"
	"path/filepath"
	"strings"
)

type GRPCServer struct {
	*Config
	UnimplementedGRPCServerServer
	internal.IEngine
}

// authenticateClient is the auth function to validate connections
func (s *GRPCServer) authenticateClient(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, fmt.Errorf("unable to extract metadata")
	}

	authStr, ok := md["authorization"]
	if !ok {
		return ctx, fmt.Errorf("authorization header missing")
	}

	if !strings.Contains(authStr[0], "Basic") {
		return ctx, fmt.Errorf("basic not found in authorization header")
	}

	decoded, err := crypto.Base64Decode(strings.Split(authStr[0], " ")[1])
	if err != nil {
		return ctx, fmt.Errorf("unable to base64 decode authorization header")
	}

	creds := strings.Split(string(decoded), ":")

	username := creds[0]
	password := creds[1]

	// check if user exists in database
	user := models.User{}
	if err = user.GetUserByUsername(s.GetSQL(), username); err != nil {
		log.Errorf(log.DatabaseLogger, "Error authenticating client, could not find user: %v", err)
		return ctx, fmt.Errorf("username/password mismatch")
	}

	// check that supplied password matches
	if password != user.Password {
		log.Errorf(log.DatabaseLogger, "Error authenticating client, invalid password supplied.")
		return ctx, fmt.Errorf("username/password mismatch")
	}

	return ctx, nil
}

// StartRPCServer starts a gRPC server with no auth
func StartRPCServer(engine internal.IEngine, config *Config) {

	log.Debugf(log.GRPClog, "gRPC server support enabled. Starting gRPC server on http://%v.\n", config.ListenAddress)
	lis, err := net.Listen("tcp", config.ListenAddress)
	if err != nil {
		log.Errorf(log.GRPClog, "gRPC server failed to bind to port: %s", err)
		return
	}

	s := GRPCServer{
		Config:  config,
		IEngine: engine,
	}

	grpcServer := grpc.NewServer()

	RegisterGRPCServerServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		//log.Fatalf("failed to serve: %s", err)
	}
}

// StartRPCServerTLS starts a gRPC server with TLS auth
func StartRPCServerTLS(engine internal.IEngine, config *Config, configDir string) *grpc.Server {
	targetDir := crypto.GetTLSDir(configDir)
	if err := crypto.CheckCerts(targetDir); err != nil {
		log.Errorf(log.GRPClog, "gRPC checkCerts failed. err: %s\n", err)
		return nil
	}
	log.Debugf(log.GRPClog, "Starting gRPC server on https://%v.\n", config.ListenAddress)
	lis, err := net.Listen("tcp", config.ListenAddress)
	if err != nil {
		log.Errorf(log.GRPClog, "gRPC server failed to bind to port: %s", err)
		return nil
	}

	creds, err := credentials.NewServerTLSFromFile(filepath.Join(targetDir, "cert.pem"), filepath.Join(targetDir, "key.pem"))
	if err != nil {
		log.Errorf(log.GRPClog, "gRPC server could not load TLS keys: %s\n", err)
		return nil
	}

	serv := GRPCServer{
		Config:  config,
		IEngine: engine,
	}
	opts := []grpc.ServerOption{
		grpc.Creds(creds),
		//grpc.UnaryInterceptor(grpcauth.UnaryServerInterceptor(serv.authenticateClient)),
	}
	server := grpc.NewServer(opts...)
	RegisterGRPCServerServer(server, &serv)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Errorf(log.GRPClog, "gRPC server failed to serve: %s\n", err)
			return
		}
	}()

	log.Debugln(log.GRPClog, "gRPC server started!")
	//if config.GRPCProxyEnabled {
	serv.StartRPCRESTProxy(configDir)
	//}
	return server

}

// StartRPCRESTProxy starts a gRPC proxy
func (s *GRPCServer) StartRPCRESTProxy(configDir string) {
	log.Debugf(log.GRPClog, "Starting gRPC proxy server on https://%v.\n", s.Config.GRPCProxyListenAddress)

	targetDir := crypto.GetTLSDir(configDir)
	creds, err := credentials.NewClientTLSFromFile(filepath.Join(targetDir, "cert.pem"), "")
	if err != nil {
		log.Errorf(log.GRPClog, "Unable to start gRPC proxy. Err: %s\n", err)
		return
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		/*grpc.WithPerRPCCredentials(auth.BasicAuth{
			Username: "admin",
			Password: "admin",
		}),*/
	}

	err = RegisterGRPCServerHandlerFromEndpoint(context.Background(), mux, s.Config.ListenAddress, opts)
	if err != nil {
		log.Errorf(log.GRPClog, "Failed to register gRPC proxy. Err: %s\n", err)
		return
	}

	server := &http.Server{
		Addr:    s.Config.GRPCProxyListenAddress,
		Handler: cors(mux),
	}

	go func() {
		if err := server.ListenAndServeTLS(filepath.Join(targetDir, "cert.pem"), filepath.Join(targetDir, "key.pem")); err != nil {
			log.Errorf(log.GRPClog, "gRPC proxy failed to serve: %s\n", err)
			return
		}
	}()

	log.Debugln(log.GRPClog, "gRPC proxy server started!")
}

// cors is a middleware function that handles settings CORS for the REST proxy.
// This code was found from https://fale.io/blog/2021/07/28/cors-headers-with-grpc-gateway
func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if allowedOrigin(r.Header.Get("Origin")) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		//}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
