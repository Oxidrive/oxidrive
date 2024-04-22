package main

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	"github.com/oxidrive/oxidrive/server/internal/config"
	"github.com/oxidrive/oxidrive/server/internal/core"
	"github.com/oxidrive/oxidrive/server/internal/web"
	"github.com/oxidrive/oxidrive/server/migrations"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)

func main() {
	trapSigterm()

	cfg := config.Parse()

	logger := InitLogger(&cfg)

	if err := migrations.Run(cfg.DatabaseConfig); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		die(logger, err, "failed to run database migrations")
	}

	db, err := sqlx.Connect(cfg.DatabaseDriver(), cfg.DatabaseSource())
	if err != nil {
		die(logger, err, "failed to open database connection")
	}

	app := core.NewApplication(core.SetupDependencies(cfg, db))

	err = web.Run(web.Config{
		Address:        cfg.ListenAddress(),
		Application:    app,
		Logger:         logger,
		FrontendFolder: cfg.AssetsFolderPath(),
	})

	if errors.Is(err, http.ErrServerClosed) {
		logger.Info().Msg("server closed")
	} else if err != nil {
		die(logger, err, "server stopped")
	}
}

func trapSigterm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()
}

func die(logger zerolog.Logger, err error, msg string) {
	logger.Error().AnErr("error", err).Msg(msg)
	os.Exit(1)
}
