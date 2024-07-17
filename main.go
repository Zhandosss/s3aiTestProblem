package main

import (
	"bankomat/configs"
	"bankomat/internal/handlers"
	"bankomat/internal/repository"
	"bankomat/internal/service"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg := configs.Load()
	log.Info().Msg("config loaded")
	log.Debug().Msgf("Config: %+v", cfg)

	e := echo.New()

	repo := repository.New()
	log.Info().Msg("repository created")

	serv := service.New(repo)
	log.Info().Msg("services created")

	handlers.New(e, serv)
	log.Info().Msg("handlers initialized")

	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      e,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	log.Info().Msgf("Server started on %s", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msgf("Server shut down: %s", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err)
	}
}
