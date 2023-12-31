package repo

import (
	"fmt"
	"klutzer/library-app/server/internal/model"
	"klutzer/library-app/shared"
	"time"

	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

// Book Repository interface
type BookRepository interface {
	Create(book *model.Book) error
	Update(id string, values map[string]interface{}) error
	Load(ids []string, author string, genre shared.Genre, publishedStart time.Time, publishedEnd time.Time) ([]model.Book, error)
	Get(id string) (model.Book, error)
}

func NewBookRepo(db *gorm.DB) BookRepository {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) Load(ids []string, author string, genre shared.Genre, publishedStart time.Time, publishedEnd time.Time) ([]model.Book, error) {

	tx := r.db.Table("book").Where("deleted = ?", false)
	if len(ids) > 0 {
		tx = tx.Where("id in ?", ids)
	}

	if author != "" {
		tx = tx.Where("author = ?", author)
	}

	if genre != "" {
		tx = tx.Where("genre = ?", genre)
	}

	if !publishedStart.IsZero() && !publishedEnd.IsZero() {
		tx = tx.Where("published_at between ? and ?", publishedStart, publishedEnd)
	} else if !publishedStart.IsZero() {
		fmt.Println("ASDASD")
		tx = tx.Where("published_at >= ?", publishedStart)
	} else if !publishedEnd.IsZero() {
		tx = tx.Where("published_at <= ?", publishedEnd)
	}

	var books []model.Book
	tx.Find(&books)

	if tx.Error != nil {
		return []model.Book{}, shared.ConvertGormErrorCode(tx.Error)
	}

	return books, nil
}

func (r *bookRepo) Create(book *model.Book) error {
	tx := r.db.Create(book)
	if tx.Error != nil {
		return shared.ConvertGormErrorCode(tx.Error)
	}

	return nil
}

func (s *bookRepo) Update(id string, values map[string]interface{}) error {
	tx := s.db.Table("book").
		Where("id = ?", id).
		Updates(values)

	if tx.Error != nil {
		return shared.ConvertGormErrorCode(tx.Error)
	}

	return nil
}

func (s *bookRepo) Get(id string) (model.Book, error) {
	var book model.Book
	tx := s.db.Table("book").
		Where("id = ?", id).
		Where("deleted = ?", false).
		First(&book)

	if tx.Error != nil {
		return model.Book{}, shared.ConvertGormErrorCode(tx.Error)
	}

	return book, nil
}
