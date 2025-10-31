package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/novrirahman-space/go-api-starter/internal/config"
	"github.com/novrirahman-space/go-api-starter/internal/logger"
	"github.com/novrirahman-space/go-api-starter/internal/server"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.Env)

	srv := server.New(cfg.HTTPAddr, log)

	// Start server async
	go func() {
		log.Info().Str("addr", cfg.HTTPAddr).Msg("starting http server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("server failed")
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Info().Msg("shutting down gracefully")
	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("graceful shutdown error")
	}
	log.Info().Msg("server stopped")
}
