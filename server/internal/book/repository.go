package book

import (
	"klutzer/conanical-library-app/shared"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

// Book Repository interface
type Repository interface {
	Create(book *Book) error
	Update(id string, values map[string]interface{}) error
	Get(id string) (Book, error)
	GetMulti(ids []string) ([]Book, error)
}

func NewRepo(db *gorm.DB) Repository {
	db.AutoMigrate(&Book{})

	return &repo{
		db: db,
	}
}

func (r *repo) Create(book *Book) error {
	tx := r.db.Create(book)
	if tx.Error != nil {
		return shared.ConvertGormErrorCode(tx.Error)
	}

	return nil
}

func (s *repo) Update(id string, values map[string]interface{}) error {
	tx := s.db.Table("book").
		Where("id = ?", id).
		Updates(values)

	if tx.Error != nil {
		return shared.ConvertGormErrorCode(tx.Error)
	}

	return nil
}

func (r *repo) Get(id string) (Book, error) {
	var book Book
	tx := r.db.First(&book, "id = ?", id).Where("deleted = ?", false)
	if tx.Error != nil {
		return Book{}, shared.ConvertGormErrorCode(tx.Error)
	}

	return book, nil
}

func (r *repo) GetMulti(ids []string) ([]Book, error) {
	if len(ids) == 0 {
		return []Book{}, nil
	}

	var book []Book
	tx := r.db.Limit(len(ids)).Where(&book, "id in ?", ids).Where("deleted = ?", false)
	if tx.Error != nil {
		return []Book{}, shared.ConvertGormErrorCode(tx.Error)
	}

	return book, nil
}
