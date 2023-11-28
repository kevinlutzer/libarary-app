package main

import (
	"errors"
	"fmt"
	"klutzer/conanical-library-app/server/internal/model"
	"klutzer/conanical-library-app/server/internal/repo"
	"klutzer/conanical-library-app/server/internal/rest"
	"klutzer/conanical-library-app/server/internal/service"
	shared "klutzer/conanical-library-app/shared"
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

// @title           Library App REST Server
// @version         1.0
// @description     This is a sample server to manage books and ocllections of books

// @contact.name   Kevin Lutzer
// @contact.email    kevinlutzer@proton.me

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	m := shared.ApiBook{}
	fmt.Printf("\n\nm %+v\n\n", m)

	// Initialize logger with the Development configuration. This formats the logs to be easily read on Stdout.
	// NewProduction() formats the logs by default in a json format meant for log aggregators
	logger, err := zap.NewDevelopment()
	if err != nil {
		os.Exit(ErrFailedZap)
	}

	logger.Info("Initializing db")
	db, err := repo.NewDB()
	if err != nil {
		logger.Error("Failed to initialize db ", zap.Error(err))
	}

	//
	// Migrate the schemas
	//

	db.Statement.Debug()

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

	logger.Info("Starting server")
	if err := r.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("Server closed", zap.Error(err))
			os.Exit(ErrServerClosed)
		}

		logger.Fatal("Server errored", zap.Error(err))
		os.Exit(ErrServerErrored)
	}
}
