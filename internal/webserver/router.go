package webserver

import (
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type Router struct {
	routes []*route
}

func (r *Router) Handler(pattern *regexp.Regexp, handler http.Handler) {
	r.routes = append(r.routes, &route{pattern, handler})
}

func (r *Router) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	r.routes = append(r.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (r *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	for _, route := range r.routes {
		if route.pattern.MatchString(request.URL.Path) {
			route.handler.ServeHTTP(response, request)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(response, request)
}
