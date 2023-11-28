package model

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

func (v *Book) ToApi() shared.ApiBook {
	d := shared.ApiBook{
		ID:          v.ID,
		Title:       v.Title,
		Author:      v.Author,
		Description: v.Description,
		Genre:       v.Genre,
		Edition:     v.Edition,
	}

	if !v.PublishedAt.IsZero() {
		d.PublishedAt = v.PublishedAt
	}

	return d
}
