package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/oxidrive/oxidrive/server/internal/application"
	"github.com/oxidrive/oxidrive/server/internal/web"
)

func main() {
	trapSigterm()

	cfg := ParseConfig()

	logger := InitLogger(&cfg)

	app := application.New()

	err := web.Run(web.Config{
		Address:        cfg.ListenAddress(),
		Application:    app,
		Logger:         logger,
		FrontendFolder: cfg.AssetsFolderPath(),
	})

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("server stopped: %s\n", err)
		os.Exit(1)
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
