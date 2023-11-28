package repo

import (
	"os"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "gorm.db"
	}

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
