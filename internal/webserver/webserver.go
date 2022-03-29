package webserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal"
	"github.com/quantstop/quantstopterminal/internal/assets"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/pkg/system/crypto"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type key int

const (
	requestIDKey key = 0
)

// Webserver is the type to access and store both http/https webserver, and ws/wss webserver.
type Webserver struct {
	*Config
	internal.IEngine
	HttpServer *http.Server
	//WebsocketServer *http.Server
	mux              *http.ServeMux
	shutdownFinished chan struct{}
}

func CreateWebserver(eng internal.IEngine, conf *Config) (*Webserver, error) {

	if eng == nil {
		return nil, errors.New("engine interface cannot be nil")
	}

	if conf == nil {
		return nil, errors.New("config cannot be nil")
	}

	// next request id used for logging
	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	// Create server
	server := &Webserver{}
	server.IEngine = eng
	server.Config = conf
	server.mux = http.NewServeMux()
	server.HttpServer = &http.Server{
		Addr:    conf.HttpListenAddr,
		Handler: tracing(nextRequestID)(logging()(server.mux)),
	}

	// return the built webserver
	return server, nil

}

func (s *Webserver) SetupRoutes(isDev bool) {

	if isDev {
		log.Debugln(log.Webserver, "Development mode: On. Starting node server ...")

	} else {
		log.Debugln(log.Webserver, "Development mode: Off. Serving static frontend.")
		s.mux.Handle("/", http.FileServer(assets.Assets))
	}

}

func (s *Webserver) ListenAndServe(tls bool, configDir string) (err error) {
	if s.shutdownFinished == nil {
		s.shutdownFinished = make(chan struct{})
	}

	if tls {
		targetDir := crypto.GetTLSDir(configDir)
		if err := crypto.CheckCerts(targetDir); err != nil {
			log.Errorf(log.GRPClog, "gRPC checkCerts failed. err: %s\n", err)
		}

		log.Debugf(log.GRPClog, "Starting webserver on https://%v.\n", s.HttpListenAddr)
		err = s.HttpServer.ListenAndServeTLS(filepath.Join(targetDir, "cert.pem"), filepath.Join(targetDir, "key.pem"))
		if err == http.ErrServerClosed {
			// expected error after calling Server.Shutdown().
			err = nil
		} else if err != nil {
			err = fmt.Errorf("unexpected error from ListenAndServe: %w", err)
			return
		}
	} else {
		log.Debugf(log.GRPClog, "Starting webserver on http://%v.\n", s.HttpListenAddr)
		err = s.HttpServer.ListenAndServe()
		if err == http.ErrServerClosed {
			// expected error after calling Server.Shutdown().
			err = nil
		} else if err != nil {
			err = fmt.Errorf("unexpected error from ListenAndServe: %w", err)
			return
		}
	}

	log.Infoln(log.Webserver, "Webserver waiting for shutdown finishing...")
	<-s.shutdownFinished
	log.Infoln(log.Webserver, "Webserver shutdown finished.")

	return
}

func (s *Webserver) Shutdown() {
	log.Infoln(log.Webserver, "Webserver is shutting down.")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := s.HttpServer.Shutdown(ctx)
	if err != nil {
		log.Errorf(log.Webserver, "Webserver could not shutdown %v\n", err)
	} else {
		log.Infoln(log.Webserver, "Webserver shutdown successful.")
		close(s.shutdownFinished)
	}
	return

}

func logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				log.Debugln(log.Webserver, requestID, r.Method, r.URL.Path, GetIP(r), r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (s *Webserver) StartNodeDevelopmentServer() {

	if s.shutdownFinished == nil {
		s.shutdownFinished = make(chan struct{})
	}

	log.Debugf(log.Webserver, "Starting node development server ...")

	var cmd *exec.Cmd
	var err error

	cmd = exec.Command("npm", "run", "serve")
	cmd.Dir = "../../web"
	cmd.Stdout = os.Stdout

	if err = cmd.Start(); err != nil {
		log.Errorf(log.Webserver, "Error starting node development server %v.\n", err)
	}

	// todo: i think this all works as intended, the problem is node runtime spawns another background process
	// todo: need to figure out how to kill all
	log.Infoln(log.Webserver, "Shutting down node development server ...")
	<-s.shutdownFinished
	if err = cmd.Process.Kill(); err != nil {
		log.Errorf(log.Webserver, "Error unable to kill process node development server %v.\n", err)
	}

	return

}
