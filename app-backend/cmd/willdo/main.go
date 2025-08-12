package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"willdo/internal/config"
	"willdo/internal/database"
	"willdo/internal/repository"
	"willdo/internal/server"
)

func main() {
	// logger setup
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)
	// env lookup
	cfg := config.Load(logger)

	// db setup
	db, err := database.InitDB(cfg.DatabaseURL(), logger)
	if err != nil {
		logger.Printf("[ERROR] Failed to initialize database: %v\n", err)
		os.Exit(1)
	}

	// event repository setup
	eventRepo := repository.NewDatabaseEventRepository(db)

	// server start
	srv := server.New(logger, cfg.ApiBaseUrl, eventRepo)

	go func() {
		if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Printf("[ERROR] Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// server shutdown
	// trap sigterm or interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// block until a signal is received
	sig := <-c
	logger.Printf("[INFO] Received signal: %v\n", sig)

	// gracefully shutdown the server with a max 30 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Printf("[ERROR] Server shutdown error: %v\n", err)
		os.Exit(1)
	}

	logger.Println("[INFO] Server gracefully stopped")
}
