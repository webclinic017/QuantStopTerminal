package grpcserver

// Good info on gRPC and auth here: https://jbrandhorst.com/post/grpc-auth/

import (
	"fmt"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/quantstop/quantstopterminal/internal"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/pkg/system/crypto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
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

	if username != s.Config.Username ||
		password != s.Config.Password {
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
func StartRPCServerTLS(engine internal.IEngine, config *Config, configDir string) {
	targetDir := crypto.GetTLSDir(configDir)
	if err := crypto.CheckCerts(targetDir); err != nil {
		log.Errorf(log.GRPClog, "gRPC checkCerts failed. err: %s\n", err)
		return
	}
	log.Debugf(log.GRPClog, "Starting gRPC server on https://%v.\n", config.ListenAddress)
	lis, err := net.Listen("tcp", config.ListenAddress)
	if err != nil {
		log.Errorf(log.GRPClog, "gRPC server failed to bind to port: %s", err)
		return
	}

	creds, err := credentials.NewServerTLSFromFile(filepath.Join(targetDir, "cert.pem"), filepath.Join(targetDir, "key.pem"))
	if err != nil {
		log.Errorf(log.GRPClog, "gRPC server could not load TLS keys: %s\n", err)
		return
	}

	serv := GRPCServer{
		Config:  config,
		IEngine: engine,
	}
	opts := []grpc.ServerOption{
		grpc.Creds(creds),
		grpc.UnaryInterceptor(grpcauth.UnaryServerInterceptor(serv.authenticateClient)),
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

	/*if s.Settings.EnableGRPCProxy {
		s.StartRPCRESTProxy()
	}*/
}

// StartRPCRESTProxy starts a gRPC proxy
func (s *GRPCServer) StartRPCRESTProxy() {
	/*log.Debugf(log.GRPClog, "Starting gRPC proxy server on http://%v.\n", s.Config.GRPCProxyListenAddress)

	targetDir := crypto.GetTLSDir(s..Config.ConfigDir)
	creds, err := credentials.NewClientTLSFromFile(filepath.Join(targetDir, "cert.pem"), "")
	if err != nil {
		log.Errorf(log.GRPClog, "Unable to start gRPC proxy. Err: %s\n", err)
		return
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(auth.BasicAuth{
			Username: s.Config.Username,
			Password: s.Config.Password,
		}),
	}
	err = grpcserver.RegisterGoCryptoTraderHandlerFromEndpoint(context.Background(),
		mux, s.Config.ListenAddress, opts)
	if err != nil {
		log.Errorf(log.GRPClog, "Failed to register gRPC proxy. Err: %s\n", err)
		return
	}

	go func() {
		if err := http.ListenAndServe(s.Config.GRPCProxyListenAddress, mux); err != nil {
			log.Errorf(log.GRPClog, "gRPC proxy failed to server: %s\n", err)
			return
		}
	}()

	log.Debugln(log.GRPClog, "gRPC proxy server started!")*/
}

// All RPC functions below -------------------------------------------------------------------------------------------

// GetInfo returns info about the current session
func (s *GRPCServer) GetInfo(_ context.Context, _ *GetInfoRequest) (*GetInfoResponse, error) {

	return &GetInfoResponse{
		Uptime:          s.GetUptime(),
		Version:         s.GetVersionString(false),
		SubsystemStatus: s.GetSubsystemsStatus(),
	}, nil
}

// GetSubsystems returns a list of subsystems and their status
func (s *GRPCServer) GetSubsystems(_ context.Context, _ *GetSubsystemsRequest) (*GetSusbsytemsResponse, error) {
	return &GetSusbsytemsResponse{SubsystemsStatus: s.GetSubsystemsStatus()}, nil
}

// EnableSubsystem enables a engine subsystem
func (s *GRPCServer) EnableSubsystem(_ context.Context, r *GenericSubsystemRequest) (*GenericResponse, error) {
	err := s.SetSubsystem(r.Subsystem, true)
	if err != nil {
		return nil, err
	}
	return &GenericResponse{Status: "success",
		Data: fmt.Sprintf("subsystem %s enabled", r.Subsystem)}, nil

}

// DisableSubsystem disables a engine subsystem
func (s *GRPCServer) DisableSubsystem(_ context.Context, r *GenericSubsystemRequest) (*GenericResponse, error) {
	err := s.SetSubsystem(r.Subsystem, false)
	if err != nil {
		return nil, err
	}
	return &GenericResponse{Status: "success",
		Data: fmt.Sprintf("subsystem %s disabled", r.Subsystem)}, nil

}
