package rest

import (
	"bytes"
	"encoding/json"
	"klutzer/conanical-library-app/server/internal/model"
	"klutzer/conanical-library-app/server/internal/repo"
	"klutzer/conanical-library-app/server/internal/service"
	"net/http"
	"os"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func makeRequest(method, url string, body interface{}) (*http.Response, error) {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	httpClient := &http.Client{}
	return httpClient.Do(request)
}

func StartServer() (Rest, repo.BookRepository, repo.CollectionRepository, string) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	dbFile := uuid.New().String() + ".db"
	db, err := repo.NewDB(dbFile)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Collection{})

	bookRepo := repo.NewBookRepo(db)
	bookService := service.NewBookService(bookRepo)

	collectionRepo := repo.NewCollectionRepo(db)
	collectionService := service.NewCollectionService(collectionRepo, bookService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := NewREST(logger, bookService, collectionService, port)

	// Start server in a goroutine so it doesn't block tests
	go func() {
		if err := r.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	return r, bookRepo, collectionRepo, dbFile

}
