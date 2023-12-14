package model

import (
	"klutzer/library-app/shared"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          string       `gorm:"primaryKey"`
	Title       string       `gorm:"type:varchar(512)"`
	Author      string       `gorm:"type:varchar(512);index"`
	Description string       `gorm:"type:varchar(4096)"`
	PublishedAt time.Time    `gorm:"index"`
	Genre       shared.Genre `gorm:"type:varchar(4096);index"`
	Edition     uint8
	Deleted     bool `gorm:"default:false;index"`
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
