package rest

import (
	"klutzer/conanical-library-app/shared"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_BookHandler_Put(t *testing.T) {
	rest, bookRepo, _, dbFile := StartServer()

	// Close the server and remove the db file
	defer func() {
		os.Remove(dbFile)
		rest.Close()
	}()

	putReq := shared.BookPutRequest{
		Title: "The Hobbit",
		Data: &shared.BookData{
			Author:      "J.R.R. Tolkien",
			Genre:       string(shared.Fantasy),
			PublishedAt: time.Date(1937, time.September, 21, 0, 0, 0, 0, time.UTC),
			Edition:     1,
		},
		ID: uuid.New().String(),
	}

	res, err := makeRequest("PUT", "http://localhost:8080/v1/book", putReq)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	book, err := bookRepo.Get(putReq.ID)
	if err != nil {
		t.Fatal(err)
	}

	// Assertions
	assert.Equal(t, putReq.ID, book.ID)
	assert.Equal(t, putReq.Title, book.Title)
	assert.Equal(t, putReq.Data.Author, book.Author)
	assert.Equal(t, putReq.Data.Description, book.Description)
	assert.Equal(t, putReq.Data.PublishedAt, book.PublishedAt)
	assert.Equal(t, putReq.Data.Genre, book.Genre)
	assert.Equal(t, putReq.Data.Edition, book.Edition)
}
