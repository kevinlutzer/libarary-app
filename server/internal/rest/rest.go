package rest

import (
	"klutzer/conanical-library-app/server/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
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
	Close() error
	ListenAndServe() error
	BookHandler(w http.ResponseWriter, r *http.Request)
	CollectionHandler(w http.ResponseWriter, r *http.Request)
}

type rest struct {
	logger            *zap.Logger
	bookService       service.BookService
	collectionService service.CollectionService
	server            *gin.Engine
}

func NewREST(logger *zap.Logger, bookService service.BookService, collectionService service.CollectionService) Rest {
	g := gin.Default()

	restServer := &rest{
		logger:            logger,
		bookService:       bookService,
		collectionService: collectionService,
		server:            g,
	}

	g.Get(apiBook, restServer.GetBookHandler)

	return restServer
}

func (restServer *rest) ListenAndServe() error {
	return restServer.server.ListenAndServe()
}

func (restServer *rest) Close() error {
	return restServer.server.Close()
}
