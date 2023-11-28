package rest

import (
	"encoding/json"
	"io"
	"klutzer/conanical-library-app/server/internal/model"
	"klutzer/conanical-library-app/shared"
	"net/http"

	"go.uber.org/zap"
)

func (restService *restService) CollectionHandler(w http.ResponseWriter, r *http.Request) {
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

		q := r.URL.Query()
		var includeBooks bool
		if q.Has("includebooks") {
			includeBooks = q.Get("includebooks") == "true"
		}

		// Log Request
		restService.logger.Info("CollectionHandler", zap.String("method", r.Method), zap.String("path", r.URL.Path))

		collections, bookMap, err := restService.collectionService.Load(includeBooks)
		if err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		apiCollection := make([]shared.ApiCollection, len(collections))
		for i := range collections {
			books := []model.Book{}
			if includeBooks {
				for _, id := range collections[i].GetIDs() {
					if val, ok := bookMap[id]; ok {
						books = append(books, val)
					}
				}
			}

			apiCollection[i] = collections[i].ToApi(books)
		}

		res := shared.CollectionLoadResponse{Collections: apiCollection}
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
		restService.logger.Info("CollectionHandler", zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("body", string(b)))
	}

	//
	// Create operation
	//
	if r.Method == http.MethodPut {
		req := shared.CollectionPutRequest{}
		if err := json.Unmarshal([]byte(b), &req); err != nil {
			err := shared.NewError(shared.InvalidArguments, err.Error())
			restService.WriteErrorResponse(w, err)
			return
		}

		if err := req.Validate(); err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		id, err := restService.collectionService.Create(req.Name, req.BookIDs)
		if err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		res := shared.CollectionPutResponse{ID: id}
		restService.WriteSuccessResponse(w, &res)
		return
	}

	//
	// Update operation
	//
	if r.Method == http.MethodPost {
		req := shared.CollectionPostRequest{}
		if err := json.Unmarshal([]byte(b), &req); err != nil {
			err := shared.NewError(shared.InvalidArguments, err.Error())
			restService.WriteErrorResponse(w, err)
			return
		}

		if err := req.Validate(); err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		restService.WriteSuccessResponse(w, nil)
		return
	}

	//
	// Delete operation
	//
	if r.Method == http.MethodDelete {
		req := shared.CollectionDeleteRequest{}
		if err := json.Unmarshal([]byte(b), &req); err != nil {
			err := shared.NewError(shared.InvalidArguments, err.Error())
			restService.WriteErrorResponse(w, err)
			return
		}

		if err := req.Validate(); err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		err := restService.collectionService.Delete(req.ID)
		if err != nil {
			restService.WriteErrorResponse(w, err)
			return
		}

		restService.WriteSuccessResponse(w, nil)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
