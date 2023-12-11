package repo

import (
	"klutzer/conanical-library-app/server/internal/model"
	"klutzer/conanical-library-app/shared"

	"gorm.io/gorm"
)

type collectionRepo struct {
	db *gorm.DB
}

// Book Repository interface
type CollectionRepository interface {
	Load() ([]model.Collection, error)
	Create(collection *model.Collection) error
	Update(id string, values map[string]interface{}) error
	Get(id string) (model.Collection, error)
}

func NewCollectionRepo(db *gorm.DB) CollectionRepository {
	return &collectionRepo{
		db: db,
	}
}

func (r *collectionRepo) Load() ([]model.Collection, error) {
	var collections []model.Collection
	tx := r.db.Table("collection").Where("deleted = ?", false).Find(&collections)
	if tx.Error != nil {
		return []model.Collection{}, shared.ConvertGormErrorCode(tx.Error)
	}

	return collections, nil
}

func (r *collectionRepo) Create(collection *model.Collection) error {
	tx := r.db.Create(collection)
	if tx.Error != nil {
		return shared.ConvertGormErrorCode(tx.Error)
	}

	return nil
}

func (s *collectionRepo) Update(id string, values map[string]interface{}) error {
	tx := s.db.Table("collection").
		Where("id = ?", id).
		Where("deleted = ?", false).
		Updates(values)

	if tx.Error != nil {
		return shared.ConvertGormErrorCode(tx.Error)
	}

	return nil
}

func (r *collectionRepo) Get(id string) (model.Collection, error) {
	var collection model.Collection
	tx := r.db.First(&collection, "id = ?", id).Where("deleted = ?", false)
	if tx.Error != nil {
		return model.Collection{}, shared.ConvertGormErrorCode(tx.Error)
	}

	return collection, nil
}
