package main

import (
	"embed"
	"flag"
	"fmt"
	"fuku/api"
	"fuku/internal"
	"io/fs"
	"net/http"
	"runtime"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port int

	//go:embed static/*
	docs embed.FS
)

func main() {
	flag.IntVar(&port, "p", 8080, "port to listen at")
	flag.Parse()

	swagger, err := fs.Sub(docs, "static")
	if err != nil {
		panic(err)
	}

	srv := newServer(port, swagger)
	srv.ListenAndServe()
}

func newServer(port int, docs fs.FS) *http.Server {
	workChan := make(chan internal.Request)

	loadBalancer := internal.NewLoadBalancer(runtime.NumCPU())
	go loadBalancer.Balance(workChan)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/convert", api.Convert(workChan))
		})
	})

	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	r.Get("/*", http.FileServer(http.FS(docs)).ServeHTTP)

	return &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}
}
