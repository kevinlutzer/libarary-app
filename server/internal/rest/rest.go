package rest

import (
	"klutzer/conanical-library-app/server/internal/service"
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

type restService struct {
	logger            *zap.Logger
	bookService       service.BookService
	collectionService service.CollectionService
}

func NewREST(logger *zap.Logger, bookService service.BookService, collectionService service.CollectionService) *http.Server {
	mux := mux.NewRouter()

	mux.StrictSlash(false)
	mux.SkipClean(true)

	restService := &restService{
		logger:            logger,
		bookService:       bookService,
		collectionService: collectionService,
	}

	mux.HandleFunc(apiBook, restService.BookHandler)
	mux.HandleFunc(apiCollection, restService.CollectionHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
