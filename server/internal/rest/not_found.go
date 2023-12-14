package rest

import (
	"klutzer/library-app/shared"

	"github.com/gin-gonic/gin"
)

func (restService *rest) NotFoundHandler(w *gin.Context) {
	restService.WriteErrorResponse(w, shared.NewError(shared.NotFound, "The specified route has no handler"))
}
