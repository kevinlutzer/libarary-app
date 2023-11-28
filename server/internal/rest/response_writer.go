package rest

import (
	"encoding/json"
	shared "klutzer/conanical-library-app/shared"
	"net/http"

	"go.uber.org/zap"
)

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

func (restService *rest) WriteErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	res := shared.ApiResponse[shared.AppError]{
		Type: "err",
	}

	if val, ok := err.(*shared.AppError); ok {
		e, _ := appErrorToHttpErrorMap[err.(*shared.AppError).Type]
		w.WriteHeader(e)
		res.Data = *val
	} else {
		res.Data = defautError
		w.WriteHeader(http.StatusInternalServerError)
	}

	b, _ := json.Marshal(res)
	w.Write(b)

	restService.logger.Error("Error response", zap.String("body", string(b)))
	return
}

func (restService *rest) WriteSuccessResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	res := shared.ApiResponse[any]{
		Type: "success",
	}

	if data != nil {
		res.Data = data
	}

	b, err := json.Marshal(res)
	if err != nil {
		restService.logger.Error("Failed to create api response", zap.Error(err))
		err := shared.NewError(shared.Internal, "Failed to create api response")
		restService.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	restService.logger.Info("Success response", zap.String("body", string(b)))
}
