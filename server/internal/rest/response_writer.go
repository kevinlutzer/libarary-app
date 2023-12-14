package rest

import (
	shared "klutzer/library-app/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
// Local Error to HTTP Error Conversion
//

var appErrorToHttpErrorMap = map[shared.ErrorType]int{
	shared.AlreadyExists:      http.StatusBadRequest,
	shared.NotFound:           http.StatusNotFound,
	shared.Internal:           http.StatusInternalServerError,
	shared.InvalidArguments:   http.StatusBadRequest,
	shared.MethodNotAllow:     http.StatusMethodNotAllowed,
	shared.PreconditionFailed: http.StatusPreconditionFailed,
}

var defautError = shared.AppError{
	Type: shared.Internal,
	Msg:  "Internal error",
}

func (restService *rest) WriteErrorResponse(w *gin.Context, err error) {

	res := shared.ApiResponse[shared.AppError]{
		Type: "err",
	}

	code := http.StatusInternalServerError
	if val, ok := err.(*shared.AppError); ok {
		e, _ := appErrorToHttpErrorMap[err.(*shared.AppError).Type]
		code = e
		res.Data = *val
	} else {
		res.Data = defautError
	}

	w.JSON(code, res)
}

func (restService *rest) WriteSuccessResponse(w *gin.Context, data any) {
	res := shared.ApiResponse[any]{
		Type: "success",
	}

	if data != nil {
		res.Data = data
	}

	w.JSON(http.StatusOK, res)
}
