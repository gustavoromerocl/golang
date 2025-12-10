package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		}
	case http.MethodPost:
		w.Write([]byte("POST method"))
	default:
		http.Error(w, "404 page", http.StatusMethodNotAllowed)
	}
}

func main() {
	api := &api{addr: ":8080"}

	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}

	srv.ListenAndServe()
}
