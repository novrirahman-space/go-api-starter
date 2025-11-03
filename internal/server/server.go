package server

import (
	"net/http"
	"time"

	"github.com/rs/cors"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/novrirahman-space/go-api-starter/internal/handlers"
	"github.com/novrirahman-space/go-api-starter/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

func New(addr string, log zerolog.Logger) *http.Server {
	r := chi.NewRouter()
	r.Use(chimw.RequestID)
	

	// Middlewares
	r.Use(middleware.RequestLogger(log))
	r.Use(middleware.Timeout(15 * time.Second))
	r.Use(middleware.RateLimit(10, 20))
	r.Use(middleware.Metrics())

	// CORS
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
		MaxAge: 300,
	}).Handler)

	// Routes
	r.Get("/health", handlers.Health)
	r.Get("/v1/example", handlers.ExampleHandler)
	r.Handle("/metrics", promhttp.Handler())
	r.Get("/healthz", handlers.Liveness)
	r.Get("/readyz", handlers.Readiness)
	
	r.Get("/docs", handlers.Redoc("/openapi.yaml"))
	r.Handle("/openapi.yaml", http.StripPrefix("/", http.FileServer(http.Dir("./api"))))

	r.Route("/v1/users", func(r chi.Router) {
		r.Get("/", handlers.ListUsers)
		r.Post("/", handlers.CreateUser)
		r.Delete("/", handlers.DeleteUser)
	})

	return &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
