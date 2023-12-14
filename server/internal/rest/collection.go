package rest

import (
	"encoding/json"
	"io"
	"klutzer/library-app/server/internal/model"
	"klutzer/library-app/shared"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Summary Get Collections
// @Param        includebooks    query     bool  false  "include the books nested as each collection, this option will increase the time the API takes to execute"
// @Description Loads all collections that are stored
// @Tags collection
// @Produce json
// @Success 200 {object} shared.CollectionGetResponse
// @Router /collection [get]
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

	res := shared.CollectionGetResponseData{Collections: apiCollection}
	restService.WriteSuccessResponse(c, &res)
}

// @BasePath /v1
// @Summary Create a Collection
// @Param request body shared.CollectionCreateRequest true "form data"
// @Description Creates a collection based on the specified data in the request body
// @Tags collection
// @Accept  json
// @Produce json
// @Success 200 {object} shared.CollectionCreateResponse
// @Router /collection [put]
func (restService *rest) CreateCollectionHandler(c *gin.Context) {
	// Read request as bytes
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(c, err)
		return
	}

	// Unmarshal request
	req := shared.CollectionCreateRequest{}
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

	res := shared.CollectionCreateResponseData{ID: id}
	restService.WriteSuccessResponse(c, &res)
}

// @BasePath /v1
// @Summary Update a Collection
// @Param request body shared.CollectionUpdateRequest true "form data"
// @Description Updates a collection with the specified id in the request body. Fields can be additionally updated based on if they appear in the field mask.
// @Tags collection
// @Accept  json
// @Produce json
// @Router /collection [post]
func (restService *rest) UpdateCollectionHandler(c *gin.Context) {
	// Read request as bytes
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(c, err)
		return
	}

	// Unmarshal request
	req := shared.CollectionUpdateRequest{}
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

// @BasePath /v1
// @Summary Delete a Collection
// @Param        id    query     string  true  "the id of the collection to delete"
// @Description Deletes a collection with the specified id, a collection will not be able to be updated, deleted or surfaced in GET /v1/collection API.
// @Produce json
// @Tags collection
// @Success 200
// @Router /collection [delete]
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
