package rest

import (
	"klutzer/conanical-library-app/server/internal/service"
	"net/http"

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

type Rest interface {
	ListenAndServe() error
	BookHandler(w http.ResponseWriter, r *http.Request)
	CollectionHandler(w http.ResponseWriter, r *http.Request)
}

type rest struct {
	logger            *zap.Logger
	bookService       service.BookService
	collectionService service.CollectionService
	server            *http.Server
}

func NewREST(logger *zap.Logger, bookService service.BookService, collectionService service.CollectionService, port string) Rest {
	mux := mux.NewRouter()

	mux.StrictSlash(false)
	mux.SkipClean(true)

	restServer := &rest{
		logger:            logger,
		bookService:       bookService,
		collectionService: collectionService,
		server: &http.Server{
			Addr:    ":" + port,
			Handler: mux,
		},
	}

	mux.HandleFunc(apiBook, restServer.BookHandler)
	mux.HandleFunc(apiCollection, restServer.CollectionHandler)

	return restServer
}

func (restServer *rest) ListenAndServe() error {
	return restServer.server.ListenAndServe()
}

func (restServer *rest) Stop() error {
	return restServer.server.Close()
}
