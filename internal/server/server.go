package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/novrirahman-space/go-api-starter/internal/handlers"
	"github.com/novrirahman-space/go-api-starter/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

func New(addr string, log zerolog.Logger) *http.Server {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestLogger(log))
	r.Use(middleware.Timeout(15 * time.Second))

	// Routes
	r.Get("/health", handlers.Health)
	r.Get("/v1/example", handlers.ExampleHandler)
	r.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
