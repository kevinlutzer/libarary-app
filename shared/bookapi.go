package shared

import (
	"slices"
	"time"
)

type BookPostRequest struct {
	ID        string    `json:"id"`
	Data      *BookData `json:"data"`
	FieldMask []string  `json:"fieldMask"`
}

func (req *BookPostRequest) Validate() error {
	if req.ID == "" {
		return NewError(InvalidArguments, "id is required")
	}

	if req.Data == nil {
		req.Data = &BookData{}
	}

	return nil
}

type BookDeleteRequest struct {
	ID string `json:"id"`
}

func (req *BookDeleteRequest) Validate() error {
	if req.ID == "" {
		return NewError(InvalidArguments, "id is required")
	}

	return nil
}

type BookData struct {
	Author      string    `json:"author"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Genre       string    `json:"genre"`
	Edition     uint8     `json:"edition"`
}

func (req *BookData) Validate() error {
	var errMsg string
	if req.Author != "" && len(req.Author) > 512 {
		errMsg = "author is required and must be less than 512 characters"
	}

	if req.Description != "" && len(req.Description) > 4046 {
		errMsg = "description is required and must be less than 4046 characters"
	}

	if req.Genre != "" && !slices.Contains(Genres, Genre(req.Genre)) {
		var genreStr string
		for _, genre := range Genres {
			genreStr += string(genre) + ", "
		}

		errMsg = "genre is not one of the valid genres: " + genreStr
	}

	if req.Edition > 255 {
		errMsg = "edition must be less then 256 required"
	}

	if errMsg != "" {
		return NewError(InvalidArguments, errMsg)
	}

	return nil
}

type BookPutRequest struct {
	Title     string    `json:"title"`
	FieldMask []string  `json:"field_mask"`
	Data      *BookData `json:"data"`
}

func (req *BookPutRequest) Validate() error {
	if req.Title == "" || len(req.Title) > 512 {
		return NewError(InvalidArguments, "title is required and must be less than 512 characters")
	}

	if req.Data == nil {
		req.Data = &BookData{}
		return nil
	}

	return req.Data.Validate()
}
