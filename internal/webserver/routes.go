package webserver

import (
	"github.com/quantstop/quantstopterminal/internal/assets"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/webserver/errors"
	"github.com/quantstop/quantstopterminal/internal/webserver/handlers"
	"github.com/quantstop/quantstopterminal/internal/webserver/router"
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
	"net/http"
	"runtime/debug"
)

func (s *Webserver) ConfigureRouter() {

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
	s.mux.FrontendHandler = http.FileServer(assets.Assets)

	s.mux.GET("/api/all", handlers.Test, router.Public)

	// Session routes
	s.mux.POST("/api/session", handlers.Login, router.Public)
	s.mux.DELETE("/api/session", handlers.Logout, router.Public)

	// Reset routes
	/*s.mux.POST("/reset", handlers.CreateReset, router.User)
	s.mux.GET("/reset/([0-9]+)", handlers.DoReset, router.User)*/

	// User routes
	s.mux.POST("/api/signup", handlers.Signup, router.Public)
	s.mux.GET("/api/user", handlers.Whoami, router.User)
	/*s.mux.PUT("/user/password", handlers.UpdatePassword, router.User)*/

}
