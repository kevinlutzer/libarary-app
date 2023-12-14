package rest

import (
	"encoding/json"
	"io"
	"klutzer/library-app/server/internal/model"
	"klutzer/library-app/shared"

	"github.com/gin-gonic/gin"
)

func (restService *rest) GetCollectionHandler(c *gin.Context) {
	var includeBooks bool
	q := c.Request.URL.Query()
	if q.Has("includebooks") {
		includeBooks = q.Get("includebooks") == "true"
	}

	// Load the collections
	collections, bookMap, err := restService.collectionService.Load(includeBooks)
	if err != nil {
		restService.WriteErrorResponse(c, err)
		return
	}

	// Convert collections to api collections
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
	restService.WriteSuccessResponse(c, &res)
}

func (restService *rest) CreateCollectionHandler(c *gin.Context) {
	// Read request as bytes
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(c, err)
		return
	}

	// Unmarshal request
	req := shared.CollectionPutRequest{}
	if err := json.Unmarshal([]byte(b), &req); err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(c, err)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		restService.WriteErrorResponse(c, err)
		return
	}

	// Create the collection
	id, err := restService.collectionService.Create(req.Name, req.BookIDs)
	if err != nil {
		restService.WriteErrorResponse(c, err)
		return
	}

	res := shared.CollectionPutResponse{ID: id}
	restService.WriteSuccessResponse(c, &res)
}

func (restService *rest) UpdateCollectionHandler(c *gin.Context) {
	// Read request as bytes
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(c, err)
		return
	}

	// Unmarshal request
	req := shared.CollectionPostRequest{}
	if err := json.Unmarshal([]byte(b), &req); err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(c, err)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		restService.WriteErrorResponse(c, err)
		return
	}

	// Update the collection
	if err := restService.collectionService.Update(req.ID, req.Data.Name, req.Data.BookIDs, req.FieldMask); err != nil {
		restService.WriteErrorResponse(c, err)
		return
	}

	restService.WriteSuccessResponse(c, nil)
}

func (restService *rest) DeleteCollectionHandler(c *gin.Context) {

	// Validate request
	if !c.Request.URL.Query().Has("id") {
		err := shared.NewError(shared.InvalidArguments, "id is required as a query param")
		restService.WriteErrorResponse(c, err)
		return
	}

	// Delete the collection
	id := c.Request.URL.Query().Get("id")
	if err := restService.collectionService.Delete(id); err != nil {
		restService.WriteErrorResponse(c, err)
		return
	}

	restService.WriteSuccessResponse(c, nil)
	return
}
