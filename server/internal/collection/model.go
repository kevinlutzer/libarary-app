package collection

import (
	"klutzer/conanical-library-app/server/internal/book"
	"time"

	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	ID        string      `gorm:"primaryKey"`
	Name      string      `gorm:"size:512"`
	Books     []book.Book `gorm:"many2many:collection_books;"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	Deleted   bool        `gorm:"default:false"`
}

func (Collection) TableName() string {
	return "collection"
}

func (v *Collection) ToApi() map[string]interface{} {
	m := make(map[string]interface{})

	m["id"] = v.ID
	m["name"] = v.Name
	m["createdAt"] = v.CreatedAt.Format(time.RFC3339)

	return m
}
