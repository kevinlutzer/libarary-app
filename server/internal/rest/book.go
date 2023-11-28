package rest

import (
	"encoding/json"
	"io"
	"klutzer/conanical-library-app/shared"
	"net/http"

	"go.uber.org/zap"
)

func (restService *restService) BookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	// Make sure we are getting the right method
	if r.Method != http.MethodGet && r.Method != http.MethodPut && r.Method != http.MethodPost && r.Method != http.MethodDelete {
		err := shared.NewError(shared.MethodNotAllow, "Method not allowed")
		restService.WriteErrorResponse(w, err)
		return
	}

	// Get operation
	if r.Method == http.MethodGet {
		req := shared.BookGetRequest{}
		req.FromQueryStr(r.URL.Query())

		// Log Request
		restService.logger.Info("BookHandler", zap.String("method", r.Method), zap.String("path", r.URL.Path))

		// Validate request
		if err := req.Validate(); err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		books, err := restService.bookService.Load(req.IDs, req.Author, shared.Genre(req.Genre), req.RangeStart, req.RangeEnd)
		if err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		apiBooks := make([]shared.ApiBook, len(books))
		for i := range books {
			apiBooks[i] = books[i].ToApi()
		}

		res := shared.BookLoadResponse{Books: apiBooks}
		restService.WriteSuccessResponse(w, &res)

		return
	}

	var err error
	var b []byte

	if r.Method == http.MethodPut || r.Method == http.MethodPost || r.Method == http.MethodDelete {
		// Read request as bytes
		b, err = io.ReadAll(r.Body)
		if err != nil {
			err := shared.NewError(shared.InvalidArguments, err.Error())
			restService.WriteErrorResponse(w, err)
			return
		}

		// Log Request
		restService.logger.Info("BookHandler", zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("body", string(b)))
	}

	// Create operation
	if r.Method == http.MethodPut {
		req := shared.BookPutRequest{}
		if err := json.Unmarshal([]byte(b), &req); err != nil {
			err := shared.NewError(shared.InvalidArguments, err.Error())
			restService.WriteErrorResponse(w, err)
			return
		}

		// Validate request
		if err := req.Validate(); err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		id, err := restService.bookService.Create(req.ID, req.Title, req.Data.Author, req.Data.Description, req.Data.PublishedAt, shared.Genre(req.Data.Genre), req.Data.Edition)
		if err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		res := shared.BookPutResponse{ID: id}
		restService.WriteSuccessResponse(w, &res)

		return
	}

	// Post operation
	if r.Method == http.MethodPost {
		req := shared.BookPostRequest{}
		if err := json.Unmarshal([]byte(b), &req); err != nil {
			err := shared.NewError(shared.InvalidArguments, err.Error())
			restService.WriteErrorResponse(w, err)
			return
		}

		// Validate request
		if err := req.Validate(); err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		err := restService.bookService.Update(req.ID, req.Data.Author, req.Data.Description, req.Data.PublishedAt, shared.Genre(req.Data.Genre), req.Data.Edition, req.FieldMask)
		if err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		restService.WriteSuccessResponse(w, nil)

		return
	}

	// Delete operation
	if r.Method == http.MethodDelete {
		req := shared.BookDeleteRequest{}
		if err := json.Unmarshal([]byte(b), &req); err != nil {
			err := shared.NewError(shared.InvalidArguments, err.Error())
			restService.WriteErrorResponse(w, err)
			return
		}

		// Validate request
		if err := req.Validate(); err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		err := restService.bookService.Delete(req.ID)
		if err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		restService.WriteSuccessResponse(w, nil)

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
