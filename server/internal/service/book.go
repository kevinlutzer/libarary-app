package service

import (
	"klutzer/library-app/server/internal/model"
	"klutzer/library-app/server/internal/repo"
	"klutzer/library-app/shared"
	"time"

	"github.com/google/uuid"
)

type bookService struct {
	repo repo.BookRepository
}

type BookService interface {
	Load(ids []string, author string, genre shared.Genre, publishedStart time.Time, publishedEnd time.Time) ([]model.Book, error)
	Create(id string, title string, author string, description string, publishedAt time.Time, genre shared.Genre, edition uint8) (string, error)
	Update(id string, author string, description string, publishedAt time.Time, genre shared.Genre, edition uint8, fieldMask []string) error
	Delete(id string) error
}

func NewBookService(repo repo.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) Load(ids []string, author string, genre shared.Genre, publishedStart time.Time, publishedEnd time.Time) ([]model.Book, error) {
	return s.repo.Load(ids, author, genre, publishedStart, publishedEnd)
}

func (s *bookService) Create(id string, title string, author string, description string, publishedAt time.Time, genre shared.Genre, edition uint8) (string, error) {
	// Generate a new UUID if one is not provided
	if id == "" {
		id = uuid.New().String()
	}

	book := &model.Book{
		Title:       title,
		ID:          id,
		Author:      author,
		Description: description,
		PublishedAt: publishedAt,
		Genre:       genre,
		Edition:     edition,
	}

	err := s.repo.Create(book)
	if err != nil {
		return "", err
	}

	return id, err
}

func (s *bookService) Delete(uuid string) error {
	_, err := s.repo.Get(uuid)
	if err != nil {
		return err
	}

	return s.repo.Update(uuid, map[string]interface{}{
		"deleted": true,
	})
}

func (s *bookService) Update(id string, author string, description string, publishedAt time.Time, genre shared.Genre, edition uint8, fieldMask []string) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	values := make(map[string]interface{})
	for _, field := range fieldMask {
		switch field {
		case "author":
			values["author"] = author
		case "description":
			values["description"] = description
		case "publishedAt":
			values["published_at"] = publishedAt
		case "genre":
			values["genre"] = genre
		case "edition":
			values["edition"] = edition
		}
	}

	return s.repo.Update(id, values)
}
