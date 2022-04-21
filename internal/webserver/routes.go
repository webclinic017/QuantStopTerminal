package webserver

import (
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/webserver/errors"
	"github.com/quantstop/quantstopterminal/internal/webserver/handlers"
	"github.com/quantstop/quantstopterminal/internal/webserver/router"
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
	"net/http"
	"runtime/debug"
)

func (s *Webserver) ConfigureRouter(isDev bool) {

	log.Debugln(log.Webserver, "Setting up middleware ... ")
	//s.router.Use(middlewares.HttpRequestLogger())

	log.Debugln(log.Webserver, "Setting up error handlers ... ")
	s.mux.MethodNotAllowed = write.Error(errors.BadRequestMethod)
	s.mux.NotFound = write.Error(errors.RouteNotFound)
	s.mux.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		log.Errorf(log.Webserver, "Panic on %s", r.URL.Path)
		debug.PrintStack()
		write.Error(errors.InternalError)(w, r)
	}

	log.Debugln(log.Webserver, "Setting up route handlers ... ")

	// Public Routes
	s.mux.GET("/api/test", handlers.Test, router.Public)
	s.mux.GET("/api/version", handlers.GetVersion, router.Public)
	s.mux.GET("/api/sub-status", handlers.GetSubsystemStatus, router.Public)
	s.mux.GET("/api/uptime", handlers.GetUptime, router.Public)

	// Session routes
	s.mux.POST("/api/session", handlers.Login, router.Public)
	s.mux.DELETE("/api/session", handlers.Logout, router.User)
	s.mux.GET("/api/refresh-token", handlers.Test, router.User)

	// Reset routes
	/*s.mux.POST("/reset", handlers.CreateReset, router.User)
	s.mux.GET("/reset/([0-9]+)", handlers.DoReset, router.User)*/

	// User routes
	s.mux.POST("/api/signup", handlers.Signup, router.Public)
	s.mux.GET("/api/user", handlers.Whoami, router.User)
	/*s.mux.PUT("/user/password", handlers.UpdatePassword, router.User)*/

	// Admin routes
	s.mux.GET("/api/get-users", handlers.GetAllUsers, router.Admin)
	s.mux.POST("/api/set-subsystem", handlers.SetSubsystem, router.Admin)
	s.mux.POST("/api/set-sysconfig", handlers.SetSystemConfig, router.Admin)

}
