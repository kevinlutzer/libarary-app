package shared

import (
	"net/url"
	"strings"
	"time"
)

type BookGetRequest struct {
	IDs        []string  `json:"ids"`
	Author     string    `json:"author"`
	Genre      string    `json:"genre"`
	RangeStart time.Time `json:"rangeStart"`
	RangeEnd   time.Time `json:"rangeEnd"`
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

	if !r.RangeStart.IsZero() {
		q.Set("rangeStart", url.QueryEscape(r.RangeStart.Format(time.RFC3339)))
	}

	if !r.RangeEnd.IsZero() {
		q.Set("rangeEnd", url.QueryEscape(r.RangeEnd.Format(time.RFC3339)))
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

	if ok := u.Has("rangeStart"); ok {
		rangeStart, _ := url.QueryUnescape(u.Get("rangeStart"))
		r.RangeStart, _ = time.Parse(time.RFC3339, rangeStart)
	}

	if ok := u.Has("rangeEnd"); ok {
		rangeEnd, _ := url.QueryUnescape(u.Get("rangeEnd"))
		r.RangeEnd, _ = time.Parse(time.RFC3339, rangeEnd)
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

	if !req.RangeStart.IsZero() && !req.RangeEnd.IsZero() && req.RangeStart.After(req.RangeEnd) {
		return NewError(InvalidArguments, "range start must be before range end")
	}

	if !req.RangeStart.IsZero() && req.RangeStart.After(time.Now()) {
		return NewError(InvalidArguments, "range start must be before now")
	}

	return nil
}

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

type BookPutRequest struct {
	Data  *BookData `json:"data"`
	Title string    `json:"title"`
}

func (req *BookPutRequest) Validate() error {
	if req.Title == "" || len(req.Title) > 512 {
		return NewError(InvalidArguments, "title is required and must be less than 512 characters")
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
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Genre       Genre     `json:"genre"`
	Edition     uint8     `json:"edition"`
}

//
// Responses
//

type BookPutResponse struct {
	ID string `json:"id"`
}

type BookLoadResponse struct {
	Books []ApiBook `json:"books"`
}
