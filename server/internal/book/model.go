package book

import (
	"klutzer/conanical-library-app/shared"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          string    `gorm:"primaryKey"`
	Title       string    `gorm:"size:512"`
	Author      string    `gorm:"size:512"`
	Description string    `gorm:"size:4046"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	PublishedAt time.Time
	Genre       shared.Genre `gorm:"size:32"`
	Edition     uint8
	Deleted     bool `gorm:"default:false"`
}

func (Book) TableName() string {
	return "book"
}

func (v *Book) ToApi() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = v.ID
	m["title"] = v.Title
	m["author"] = v.Author
	m["description"] = v.Description

	if !v.PublishedAt.IsZero() {
		m["publishedAt"] = v.PublishedAt.Format(time.RFC3339)
	}

	m["genre"] = v.Genre
	m["edition"] = v.Edition

	return m
}
