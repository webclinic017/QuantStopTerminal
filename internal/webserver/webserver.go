package webserver

import (
	"context"
	"embed"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal"
	"github.com/quantstop/quantstopterminal/internal/assets"
	"github.com/quantstop/quantstopterminal/internal/log"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
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
	HttpServer      *http.Server
	WebsocketServer *http.Server
	mux             *http.ServeMux
}

func CreateWebserver(eng internal.IEngine, conf *Config, tls bool) (*Webserver, error) {

	var err error

	if tls {
		// Create https:// server
		// return the built webserver
		return nil, nil

	} else {
		// Create http:// server
		server := &Webserver{}
		server.IEngine = eng
		server.Config = conf
		server.mux = http.NewServeMux()
		server.HttpServer, err = createHttpServer(server.Config.HttpListenAddr, server.mux)
		if err != nil {
			return nil, err
		}

		// return the built webserver
		return server, nil
	}

}

func createHttpServer(addr string, handler http.Handler) (*http.Server, error) {

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	return &http.Server{
		Addr:     addr,
		Handler:  tracing(nextRequestID)(logging()(handler)),
		ErrorLog: &log.QSTLogger.Logger,
	}, nil
}

func (s *Webserver) SetupRoutes(isDev bool) {

	//s.HttpLogHandler = LogHandler{}
	th := TimeHandler{Format: time.RFC1123}
	vh := VersionHandler{Version: s.GetVersionString(false)}

	if isDev {
		log.Debugln(log.Webserver, "Development mode: On. Only serving api routes.")

	} else {
		log.Debugln(log.Webserver, "Development mode: Off. Serving static assets.")
		s.mux.Handle("/", http.FileServer(assets.Assets))
	}

	//s.mux.Handle("/", http.FileServer(GetWebFrontend(isDev, *Website)))
	s.mux.Handle("/api/time", th)
	s.mux.Handle("/api/version", vh)
}

func (s *Webserver) StartWebServer(isDev bool, shutdown chan struct{}) {

	// Start the Node client app (only for version "development")
	if isDev {
		StartNodeDevelopmentServer()
	}

	done := make(chan bool)

	go func() {
		<-shutdown
		log.Infoln(log.Webserver, "Webserver is shutting down.")
		//atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		s.HttpServer.SetKeepAlivesEnabled(false)
		if err := s.HttpServer.Shutdown(ctx); err != nil {
			log.Errorf(log.Webserver, "Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Infof(log.Webserver, "Starting web server, listening on http://%v\n", s.Config.HttpListenAddr)
	if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// unexpected error. port in use?
		log.Errorf(log.Webserver, "Could not listen on %s: %v\n", s.Config.HttpListenAddr, err)
	}

	<-done
	log.Infoln(log.Webserver, "Webserver stopped.")

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

func StartNodeDevelopmentServer() {

	go func() {
		log.Debugf(log.Webserver, "Starting node development server ...")

		var cmd *exec.Cmd
		var err error

		cmd = exec.Command("npm", "run", "serve")
		cmd.Dir = "../../web"
		cmd.Stdout = os.Stdout

		if err = cmd.Start(); err != nil {
			log.Errorf(log.Webserver, "Error starting node development server %v.\n", err)
		}

		// Wait for command to stop running, ie. node server is stopped
		_ = cmd.Wait()

		log.Infoln(log.Webserver, "Shutting down node development server ...")

	}()

}

// GetWebFrontend returns http.Filesystem implementation of a web frontend.
// devMode determines if that is as an embedded pointer or not.
func GetWebFrontend(devMode bool, embedFS embed.FS) http.FileSystem {

	// If in development mode, just serve the directory on the local disk.
	if devMode {
		log.Debugln(log.Webserver, "Development mode: On. Using directory on disk.")
		return http.FS(os.DirFS("assets/*"))
	}

	// If not in development mode, use embedded filesystem
	log.Debugln(log.Webserver, "Development mode: Off. Using embedded filesystem.")
	fsys, err := fs.Sub(embedFS, "assets/*")
	if err != nil {
		log.Errorf(log.Webserver, "Unable to start webserver. Error getting filesystem: %v\n", err)
	}
	return http.FS(fsys)
}
