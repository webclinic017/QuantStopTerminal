package webserver

import (
	"net/http"
	"time"
)

type IndexHandler struct {
	Assets http.FileSystem
}

func (ih IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	handler := http.FileServer(ih.Assets)
	handler.ServeHTTP(w, r)

	w.Write([]byte(""))
}

type TimeHandler struct {
	Format string
}

func (th TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	tm := time.Now().Format(th.Format)
	w.Write([]byte("The time is: " + tm))
}

type VersionHandler struct {
	Version string
}

func (v VersionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Write([]byte(v.Version))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
