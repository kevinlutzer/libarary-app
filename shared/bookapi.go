package shared

import (
	"net/url"
	"strings"
	"time"
)

type BookGetRequest struct {
	IDs            []string  `json:"ids"`
	Author         string    `json:"author"`
	Genre          string    `json:"genre"`
	PublishedStart time.Time `json:"publishedStart"`
	PublishedEnd   time.Time `json:"publishedEnd"`
}

func (r *BookGetRequest) ToQueryStr() string {
	q := url.Values{}

	if len(r.IDs) > 0 {
		q.Set("ids", url.QueryEscape(strings.Join(r.IDs, ",")))
	}

	if r.Author != "" {
		q.Set("author", url.QueryEscape(r.Author))
	}

	if r.Genre != "" {
		q.Set("genre", url.QueryEscape(r.Genre))
	}

	if !r.PublishedStart.IsZero() {
		q.Set("publishedStart", url.QueryEscape(r.PublishedStart.Format(time.RFC3339)))
	}

	if !r.PublishedEnd.IsZero() {
		q.Set("publishedEnd", url.QueryEscape(r.PublishedEnd.Format(time.RFC3339)))
	}

	return q.Encode()
}

func (r *BookGetRequest) FromQueryStr(u url.Values) {
	if ok := u.Has("ids"); ok {
		ids, _ := url.QueryUnescape(u.Get("ids"))
		r.IDs = strings.Split(ids, ",")
	}

	if ok := u.Has("author"); ok {
		author, _ := url.QueryUnescape(u.Get("author"))
		r.Author = author
	}

	if ok := u.Has("genre"); ok {
		genre, _ := url.QueryUnescape(u.Get("genre"))
		r.Genre = genre
	}

	if ok := u.Has("publishedStart"); ok {
		publishedStart, _ := url.QueryUnescape(u.Get("publishedStart"))
		r.PublishedStart, _ = time.Parse(time.DateOnly, publishedStart)
	}

	if ok := u.Has("publishedEnd"); ok {
		publishedEnd, _ := url.QueryUnescape(u.Get("publishedEnd"))
		r.PublishedEnd, _ = time.Parse(time.DateOnly, publishedEnd)
	}
}

func (req *BookGetRequest) Validate() error {
	if len(req.IDs) > 0 {
		for _, id := range req.IDs {
			if !IsValidID(id) {
				return NewError(InvalidArguments, "an id specified is not a valid id")
			}
		}
	}

	if req.Author != "" && len(req.Author) > 512 {
		return NewError(InvalidArguments, "author must be less than 512 characters")
	}

	if req.Genre != "" && !IsValidGenre(req.Genre) {
		return NewError(InvalidArguments, "genre is not one of the valid genres: "+ValidGenreStr)
	}

	if !req.PublishedStart.IsZero() && !req.PublishedEnd.IsZero() && req.PublishedStart.After(req.PublishedEnd) {
		return NewError(InvalidArguments, "range start must be before range end")
	}

	if !req.PublishedStart.IsZero() && req.PublishedStart.After(time.Now()) {
		return NewError(InvalidArguments, "range start must be before now")
	}

	return nil
}

type BookUpdateRequest struct {
	ID        string    `json:"id" binding:"required"`
	Data      *BookData `json:"data"`
	FieldMask []string  `json:"fieldMask"`
}

func (req *BookUpdateRequest) Validate() error {
	if req.ID == "" {
		return NewError(InvalidArguments, "id is required")
	}

	if req.Data == nil {
		req.Data = &BookData{}
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

	if !IsValidGenre(req.Genre) {
		errMsg = "genre is not one of the valid genres: " + ValidGenreStr
	}

	if req.Edition > 255 {
		errMsg = "edition must be less then 256 required"
	}

	if errMsg != "" {
		return NewError(InvalidArguments, errMsg)
	}

	return nil
}

type BookCreateRequest struct {
	ID    string    `json:"id"`                       // optional
	Data  *BookData `json:"data"`                     // optional
	Title string    `json:"title" binding:"required"` // required
}

func (req *BookCreateRequest) Validate() error {
	if req.Title == "" || len(req.Title) > 512 {
		return NewError(InvalidArguments, "title is required and must be less than 512 characters")
	}

	if req.ID != "" {
		if ok := IsValidID(req.ID); !ok {
			return NewError(InvalidArguments, InvalidIdMsg)
		}
	}

	if req.Data != nil {
		return req.Data.Validate()
	}

	return nil
}

//
// ApiBook Definition
//

type ApiBook struct {
	ID          string    `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Description string    `json:"description" binding:"required"`
	PublishedAt time.Time `json:"publishedAt" binding:"required"`
	Genre       Genre     `json:"genre" binding:"required"`
	Edition     uint8     `json:"edition" binding:"required"`
}

//
// Responses
//

type BookCreateResponseData struct {
	ID string `json:"id" binding:"required"`
}

type BookCreateResponse = ApiResponse[BookCreateResponseData]

type BookGetResponseData struct {
	Books []ApiBook `json:"books" binding:"required"`
}

type BookGetResponse = ApiResponse[BookGetResponseData]
