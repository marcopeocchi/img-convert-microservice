package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "p", 8080, "port to listen at")
	flag.Parse()

	srv := newServer(port)
	srv.ListenAndServe()
}

func newServer(port int) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/convert", Convert)
		})
	})

	r.Get("/", http.FileServer(http.Dir("static")).ServeHTTP)

	return &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}
}
