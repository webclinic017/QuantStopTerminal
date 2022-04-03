package middleware

import (
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
	"net/http"
	"os"
)

// Cors adds CORS headers to the response
func Cors(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//log.Println("Middleware chain | 4 | Cors")
		/*if skipCorsAndCSRF(r.URL.Path) {
			fn(w, r)
			return
		}

		if origin := validateOrigin(r); origin != "" {
			// if we were given an origin that matches our list
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

		if r.Method == http.MethodOptions {
			// simple response for the preflight check
			fn = write.Success()
		}*/

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == http.MethodOptions {
			//http.Error(w, "No Content", http.StatusNoContent)
			fn = write.Success()
			//return
		}

		fn(w, r)
	}
}

const localDev = "https://localhost"

// only returns an origin if it matches our list
func validateOrigin(r *http.Request) string {
	origin := r.Header.Get("Origin")
	apiRoot := os.Getenv("API_ROOT")
	switch origin {
	case apiRoot:
		return apiRoot
	case localDev:
		return localDev
	default:
		return ""
	}
}

// a list of paths to bypass cors checks - this is useful for webhooks and stuff
var bypassPaths = []string{
	"/",
}

func skipCorsAndCSRF(path string) bool {
	for _, c := range bypassPaths {
		if path == c {
			return true
		}
	}

	return false
}
