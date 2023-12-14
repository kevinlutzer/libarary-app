package rest

import (
	"encoding/json"
	"io"
	shared "klutzer/library-app/shared"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1
// @Summary Get Books
// @Param        ids    query     []string  false  "a list of ids of books"
// @Param        author    query     string  false  "the author of the books"
// @Param        genre    query     shared.Genre  false  "the genre of the books"
// @Param        publishedStart    query     string  false  "the start of the range of published dates must be specified in the form of 2006-01-02"
// @Param        publishedEnd    query     string  false  "the end of the range of published dates must be specified in the form of 2006-01-02"
// @Description Loads a list of books based on the specified filters in the query string
// @Tags book
// @Produce json
// @Success 200 {object} shared.BookGetResponse
// @Router /book [get]
func (restService *rest) GetBookHandler(r *gin.Context) {
	req := shared.BookGetRequest{}
	req.FromQueryStr(r.Request.URL.Query())

	// Validate request
	if err := req.Validate(); err != nil {
		restService.WriteErrorResponse(r, err)
		return
	}

	books, err := restService.bookService.Load(req.IDs, req.Author, shared.Genre(req.Genre), req.PublishedStart, req.PublishedEnd)
	if err != nil {
		restService.WriteErrorResponse(r, err)
		return
	}

	apiBooks := make([]shared.ApiBook, len(books))
	for i := range books {
		apiBooks[i] = books[i].ToApi()
	}

	res := shared.BookGetResponseData{Books: apiBooks}
	restService.WriteSuccessResponse(r, &res)

	return
}

// @BasePath /v1
// @Summary Create a Book
// @Param request body shared.BookCreateRequest true "form data"
// @Description Creates a book based on the specified data in the request body
// @Tags book
// @Accept  json
// @Produce json
// @Success 200 {object} shared.BookCreateResponse
// @Router /book [put]
func (restService *rest) CreateBookHandler(r *gin.Context) {
	// Read request as bytes
	b, err := io.ReadAll(r.Request.Body)
	if err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(r, err)
		return
	}

	// Create operation
	req := shared.BookCreateRequest{}
	if err := json.Unmarshal([]byte(b), &req); err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(r, err)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		restService.WriteErrorResponse(r, err)
		return
	}

	id, err := restService.bookService.Create(req.ID, req.Title, req.Data.Author, req.Data.Description, req.Data.PublishedAt, shared.Genre(req.Data.Genre), req.Data.Edition)
	if err != nil {
		restService.WriteErrorResponse(r, err)
		return
	}

	res := shared.BookCreateResponseData{ID: id}
	restService.WriteSuccessResponse(r, &res)
}

// @BasePath /v1
// @Summary Update a Book
// @Param request body shared.BookUpdateRequest true "form data"
// @Description Updates a book with the specified id in the request body. Fields can be additionally updated based on if they appear in the field mask.
// @Accept  json
// @Tags book
// @Produce json
// @Router /book [post]
func (restService *rest) UpdateBookHandler(r *gin.Context) {
	// Read request as bytes
	b, err := io.ReadAll(r.Request.Body)
	if err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(r, err)
		return
	}

	// Post operation
	req := shared.BookUpdateRequest{}
	if err := json.Unmarshal([]byte(b), &req); err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(r, err)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		restService.WriteErrorResponse(r, err)
		return
	}

	err = restService.bookService.Update(req.ID, req.Data.Author, req.Data.Description, req.Data.PublishedAt, shared.Genre(req.Data.Genre), req.Data.Edition, req.FieldMask)
	if err != nil {
		restService.WriteErrorResponse(r, err)
		return
	}

	restService.WriteSuccessResponse(r, nil)
}

// @BasePath /v1
// @Summary Delete a Book
// @Param        id    query     string  true  "the id of the book to delete"
// @Description Deletes a book with the specified id, a deleted book will not be able to be updated, deleted or surfaced in GET /v1/book and GET /v1/collection APIs.
// @Produce json
// @Success 200
// @Tags book
// @Router /book [delete]
func (restService *rest) DeleteBookHandler(r *gin.Context) {

	// Validate the request
	if !r.Request.URL.Query().Has("id") {
		err := shared.NewError(shared.InvalidArguments, "id is required as a query param")
		restService.WriteErrorResponse(r, err)
		return
	}

	// Delete the book
	id := r.Request.URL.Query().Get("id")
	if err := restService.bookService.Delete(id); err != nil {
		restService.WriteErrorResponse(r, err)
		return
	}

	restService.WriteSuccessResponse(r, nil)
}
