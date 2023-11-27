package book

import (
	"klutzer/conanical-library-app/shared"
	"time"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

type Service interface {
	GetMulti(ids []string) ([]Book, error)
	Create(title string, author string, description string, publishedAt time.Time, genre shared.Genre, edition uint8, fieldMask []string) (string, error)
	Update(id string, author string, description string, publishedAt time.Time, genre shared.Genre, edition uint8, fieldMask []string) error
	Delete(id string) error
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(title string, author string, description string, publishedAt time.Time, genre shared.Genre, edition uint8, fieldMask []string) (string, error) {
	id := uuid.New().String()
	book := &Book{
		Title: title,
		ID:    id,
	}

	for _, field := range fieldMask {
		switch field {
		case "author":
			book.Author = author
		case "description":
			book.Description = description
		case "published_at":
			book.PublishedAt = publishedAt
		case "genre":
			book.Genre = genre
		case "edition":
			book.Edition = edition
		}
	}

	err := s.repo.Create(book)
	if err != nil {
		return "", err
	}

	return id, err
}

func (s *service) Delete(uuid string) error {
	_, err := s.repo.Get(uuid)
	if err != nil {
		return err
	}

	return s.repo.Update(uuid, map[string]interface{}{
		"deleted": true,
	})
}

func (s *service) Update(id string, author string, description string, publishedAt time.Time, genre shared.Genre, edition uint8, fieldMask []string) error {
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

func (s *service) GetMulti(ids []string) ([]Book, error) {
	return s.repo.GetMulti(ids)
}
