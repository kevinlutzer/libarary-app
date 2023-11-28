package model

import (
	"klutzer/conanical-library-app/shared"
	"strings"

	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	Name    string `gorm:"type:varchar(512)"`
	BookIDs string `gorm:"type:text"`
	Deleted bool   `gorm:"default:false"`
}

func (v *Collection) GetIDs() []string {
	if v.BookIDs == "" {
		return []string{}
	}

	return strings.Split(v.BookIDs, ",")
}

func (Collection) TableName() string {
	return "collection"
}

func (v *Collection) ToApi(Books []Book) shared.ApiCollection {

	m := shared.ApiCollection{
		ID:    v.ID,
		Name:  v.Name,
		Books: []shared.ApiBook{},
	}

	// Add books if they are present
	if len(Books) > 0 {
		m.Books = make([]shared.ApiBook, len(Books))
		for i, book := range Books {
			m.Books[i] = book.ToApi()
		}
	}

	return m
}
