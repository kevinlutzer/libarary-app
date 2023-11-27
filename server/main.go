package main

import (
	"errors"
	"klutzer/conanical-library-app/server/internal/book"
	"klutzer/conanical-library-app/server/internal/collection"
	"net/http"
	"os"

	"go.uber.org/zap"
)

const (
	ErrDBInitFailed  = 4
	ErrFailedZap     = 5
	ErrServerClosed  = 6
	ErrServerErrored = 7
)

func main() {

	// Initialize logger with the Development configuration. This formats the logs to be easily read on Stdout.
	// NewProduction() formats the logs by default in a json format meant for log aggregators
	logger, err := zap.NewDevelopment()
	if err != nil {
		os.Exit(ErrFailedZap)
	}

	logger.Info("Initializing db")
	db, err := NewDB()
	if err != nil {
		logger.Error("Failed to initialize db ", zap.Error(err))
	}

	bookRepo := book.NewRepo(db)
	bookService := book.NewService(bookRepo)

	collectionRepo := collection.NewRepo(db)
	collectionService := collection.NewService(collectionRepo, bookService)

	server := NewREST(logger, bookService, collectionService)

	logger.Info("Starting server")
	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("Server closed", zap.Error(err))
			os.Exit(ErrServerClosed)
		}

		logger.Fatal("Server errored", zap.Error(err))
		os.Exit(ErrServerErrored)
	}
}
