package main

import (
	"fmt"
	"klutzer/library-app/server/internal/model"
	"klutzer/library-app/server/internal/repo"
	"klutzer/library-app/server/internal/rest"
	"klutzer/library-app/server/internal/service"
	"os"
)

const (
	ErrDBInitFailed = 4
	ErrServerClosed = 6
)

// @title           Library App API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Kevin Lutzer
// @contact.email  kevinlutzer@proton.me

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {

	fmt.Println("Initializing db")

	//
	// Initialize db
	//

	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "gorm.db"
	}

	db, err := repo.NewDB(dbFile)
	if err != nil {
		os.Exit(ErrDBInitFailed)
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

	fmt.Println("Initializing server")
	r := rest.NewREST(bookService, collectionService)
	if err := r.Run(":" + port); err != nil {
		os.Exit(ErrServerClosed)
	}
}
