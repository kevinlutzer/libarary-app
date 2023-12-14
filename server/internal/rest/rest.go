package rest

import (
	"klutzer/library-app/docs"
	"klutzer/library-app/server/internal/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	apiBook       = "/v1/book"
	apiCollection = "/v1/collection"
)

var apis = []string{
	apiBook, apiCollection,
}

type Rest interface {
	Run(port string) error
}

type rest struct {
	bookService       service.BookService
	collectionService service.CollectionService
	server            *gin.Engine
}

func NewREST(bookService service.BookService, collectionService service.CollectionService) Rest {
	g := gin.Default()

	restServer := &rest{
		bookService:       bookService,
		collectionService: collectionService,
		server:            g,
	}

	g.SetTrustedProxies(nil)

	// Not Found Handler
	g.NoRoute(restServer.NotFoundHandler)

	// Swagger APIs
	docs.SwaggerInfo.BasePath = "/v1"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Book APIs
	g.GET(apiBook, restServer.GetBookHandler)
	g.PUT(apiBook, restServer.CreateBookHandler)
	g.POST(apiBook, restServer.UpdateBookHandler)
	g.DELETE(apiBook, restServer.DeleteBookHandler)

	// Collection APIs
	g.GET(apiCollection, restServer.GetCollectionHandler)
	g.PUT(apiCollection, restServer.CreateCollectionHandler)
	g.POST(apiCollection, restServer.UpdateCollectionHandler)
	g.DELETE(apiCollection, restServer.DeleteCollectionHandler)

	return restServer
}

func (restServer *rest) Run(port string) error {
	return restServer.server.Run(port)
}
