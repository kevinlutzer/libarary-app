package collection

import (
	"fmt"
	"klutzer/conanical-library-app/server/internal/book"

	"github.com/google/uuid"
)

type service struct {
	repo        Repository
	bookService book.Service
}

type Service interface {
	Create(name string, bookIDs []string) (string, error)
	// Update(id string, name string, bookIDs []string, fieldMask []string) error
	Delete(id string) error
}

func NewService(repo Repository, bookService book.Service) Service {
	return &service{repo: repo, bookService: bookService}
}

func (s *service) Create(name string, bookIDs []string) (string, error) {
	books, err := s.bookService.GetMulti(bookIDs)
	if err != nil {
		return "", err
	}

	fmt.Println(books)

	id := uuid.New().String()
	err = s.repo.Create(&Collection{
		Name:  name,
		ID:    id,
		Books: books,
	})

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

// func (s *service) Update(id string, name string, bookIDs []string, fieldMask []string) error {
// 	return nil
// }
