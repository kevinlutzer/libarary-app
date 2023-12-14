package service

import (
	"fmt"
	"klutzer/library-app/server/internal/model"
	"klutzer/library-app/server/internal/repo"
	"klutzer/library-app/shared"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

type collectionService struct {
	repo        repo.CollectionRepository
	bookService BookService
}

type CollectionService interface {
	Load(includeBooks bool) ([]model.Collection, map[string]model.Book, error)
	Create(name string, bookIDs []string) (string, error)
	Update(id string, name string, bookIDs []string, fieldMask []string) error
	Delete(id string) error
}

func NewCollectionService(repo repo.CollectionRepository, bookService BookService) CollectionService {
	return &collectionService{repo: repo, bookService: bookService}
}

func (s *collectionService) Load(includeBooks bool) ([]model.Collection, map[string]model.Book, error) {

	bookMaps := map[string]model.Book{}
	collections, err := s.repo.Load()
	if err != nil {
		return []model.Collection{}, bookMaps, err
	}

	if includeBooks {
		// Collect all book IDs
		bookIDs := []string{}
		for _, collection := range collections {
			for _, id := range collection.GetIDs() {
				bookIDs = append(bookIDs, id)
			}
		}

		if len(bookIDs) > 0 {
			// Load the books we need
			books, err := s.bookService.Load(bookIDs, "", "", time.Time{}, time.Time{})
			if err != nil {
				return []model.Collection{}, bookMaps, nil
			}

			fmt.Printf("\n\nbooks %+v\n\n", books)

			// Create a map of books for easy lookup when forming the api structs
			for _, book := range books {
				bookMaps[book.ID] = book
			}
		}
	}

	return collections, bookMaps, nil
}

func (s *collectionService) hasBooks(bookIDs []string) error {
	books, err := s.bookService.Load(bookIDs, "", "", time.Time{}, time.Time{})
	if err != nil {
		return err
	}

	//
	// Check that all books exist in the collection to create
	//
	if len(books) > 0 {
		m := map[string]string{}
		for _, bookID := range bookIDs {
			m[bookID] = bookID
		}

		for _, book := range books {
			if _, ok := m[book.ID]; !ok {
				return shared.NewError(shared.PreconditionFailed, "book id "+book.ID+" does not exist")
			}
		}
	}

	return nil
}

func (s *collectionService) Create(name string, bookIDs []string) (string, error) {
	if err := s.hasBooks(bookIDs); err != nil {
		return "", err
	}

	id := uuid.New().String()
	m := &model.Collection{
		Name:    name,
		ID:      id,
		BookIDs: strings.Join(bookIDs, ","),
	}
	if err := s.repo.Create(m); err != nil {
		return "", err
	}

	return id, nil
}

func (s *collectionService) Update(id string, name string, bookIDs []string, fieldMask []string) error {
	if _, err := s.repo.Get(id); err != nil {
		return err
	}

	values := map[string]interface{}{}
	if slices.Contains(fieldMask, "bookIDs") {
		if err := s.hasBooks(bookIDs); err != nil {
			return err
		}

		values["book_ids"] = strings.Join(bookIDs, ",")
	}

	if slices.Contains(fieldMask, "name") {
		values["name"] = name
	}

	return s.repo.Update(id, values)
}

func (s *collectionService) Delete(uuid string) error {
	_, err := s.repo.Get(uuid)
	if err != nil {
		return err
	}

	return s.repo.Update(uuid, map[string]interface{}{
		"deleted": true,
	})
}
