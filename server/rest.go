package main

import (
	"klutzer/conanical-library-app/server/internal/book"
	"klutzer/conanical-library-app/server/internal/collection"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	apiBook       = "/v1/book"
	apiCollection = "/v1/collection"
)

var apis = []string{
	apiBook, apiCollection,
}

func NewREST(logger *zap.Logger, bookService book.Service, collectionService collection.Service) *http.Server {
	mux := mux.NewRouter()

	mux.StrictSlash(false)
	mux.SkipClean(true)

	mux.HandleFunc(apiBook, func(w http.ResponseWriter, r *http.Request) {
		book.BookHandler(logger, bookService, w, r)
	})

	mux.HandleFunc(apiCollection, func(w http.ResponseWriter, r *http.Request) {
		collection.CollectionHandler(logger, collectionService, w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
