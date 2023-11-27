package collection

import (
	"klutzer/conanical-library-app/shared"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

// Book Repository interface
type Repository interface {
	Create(collection *Collection) error
	Update(id string, values map[string]interface{}) error
	Get(id string) (Collection, error)
}

func NewRepo(db *gorm.DB) Repository {
	db.AutoMigrate(&Collection{})

	return &repo{
		db: db,
	}
}

func (r *repo) Create(collection *Collection) error {
	tx := r.db.Create(collection)
	if tx.Error != nil {
		return shared.ConvertGormErrorCode(tx.Error)
	}

	return nil
}

func (s *repo) Update(id string, values map[string]interface{}) error {
	tx := s.db.Table("collection").
		Where("id = ?", id).
		Updates(values)

	if tx.Error != nil {
		return shared.ConvertGormErrorCode(tx.Error)
	}

	return nil
}

func (r *repo) Get(id string) (Collection, error) {
	var collection Collection
	tx := r.db.First(&collection, "id = ?", id).Where("deleted = ?", false)
	if tx.Error != nil {
		return Collection{}, shared.ConvertGormErrorCode(tx.Error)
	}

	return collection, nil
}
