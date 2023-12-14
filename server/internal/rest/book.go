package rest

import (
	"encoding/json"
	"io"
	shared "klutzer/library-app/shared"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /v1/book [get]
func (restService *rest) GetBookHandler(r *gin.Context) {
	req := shared.BookGetRequest{}
	req.FromQueryStr(r.Request.URL.Query())

	// Validate request
	if err := req.Validate(); err != nil {
		restService.WriteErrorResponse(r, err)
		return
	}

	books, err := restService.bookService.Load(req.IDs, req.Author, shared.Genre(req.Genre), req.RangeStart, req.RangeEnd)
	if err != nil {
		restService.WriteErrorResponse(r, err)
		return
	}

	apiBooks := make([]shared.ApiBook, len(books))
	for i := range books {
		apiBooks[i] = books[i].ToApi()
	}

	res := shared.BookLoadResponse{Books: apiBooks}
	restService.WriteSuccessResponse(r, &res)

	return
}

func (restService *rest) CreateBookHandler(r *gin.Context) {
	// Read request as bytes
	b, err := io.ReadAll(r.Request.Body)
	if err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(r, err)
		return
	}

	// Create operation
	req := shared.BookPutRequest{}
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

	res := shared.BookPutResponse{ID: id}
	restService.WriteSuccessResponse(r, &res)
}

func (restService *rest) UpdateBookHandler(r *gin.Context) {
	// Read request as bytes
	b, err := io.ReadAll(r.Request.Body)
	if err != nil {
		err := shared.NewError(shared.InvalidArguments, err.Error())
		restService.WriteErrorResponse(r, err)
		return
	}

	// Post operation
	req := shared.BookPostRequest{}
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
