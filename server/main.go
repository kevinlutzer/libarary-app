package main

import (
	"klutzer/conanical-library-app/server/internal/model"
	"klutzer/conanical-library-app/server/internal/repo"
	"klutzer/conanical-library-app/server/internal/rest"
	"klutzer/conanical-library-app/server/internal/service"
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

	//
	// Initialize db
	//

	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "gorm.db"
	}

	db, err := repo.NewDB(dbFile)
	if err != nil {
		logger.Error("Failed to initialize db ", zap.Error(err))
	}

	//
	// Migrate the schemas
	//

	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Collection{})

	//
	// Initialize services, server and repo
	//

	bookRepo := repo.NewBookRepo(db)
	bookService := service.NewBookService(bookRepo)

	collectionRepo := repo.NewCollectionRepo(db)
	collectionService := service.NewCollectionService(collectionRepo, bookService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := rest.NewREST(logger, bookService, collectionService, port)
	if err := r.Run(":" + port); err != nil {
		logger.Fatal("Server closed", zap.Error(err))
		os.Exit(ErrServerClosed)
	}
}
