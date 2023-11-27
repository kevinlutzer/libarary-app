package book

import (
	"encoding/json"
	"io"
	"klutzer/conanical-library-app/shared"
	"net/http"

	"go.uber.org/zap"
)

func BookHandler(logger *zap.Logger, service Service, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	// Read request as bytes
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	logger.Info("BookHandler", zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("body", string(b)))

	// Create operation
	if r.Method == http.MethodPut {
		req := shared.BookPutRequest{}
		if err := json.Unmarshal([]byte(b), &req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Validate request
		if err := req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		_, err := service.Create(req.Title, req.Data.Author, req.Data.Description, req.Data.PublishedAt, shared.Genre(req.Data.Genre), req.Data.Edition, req.FieldMask)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		return
	}

	// Post operation
	if r.Method == http.MethodPost {
		req := shared.BookPostRequest{}
		if err := json.Unmarshal([]byte(b), &req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Validate request
		if err := req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		err := service.Update(req.ID, req.Data.Author, req.Data.Description, req.Data.PublishedAt, shared.Genre(req.Data.Genre), req.Data.Edition, req.FieldMask)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		return
	}

	// Delete operation
	if r.Method == http.MethodDelete {
		req := shared.BookDeleteRequest{}
		if err := json.Unmarshal([]byte(b), &req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Validate request
		if err := req.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		err := service.Delete(req.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
