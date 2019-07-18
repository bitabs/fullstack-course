package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

type DB struct {
	*gorm.DB
}

func (db *DB) Connect() (*DB, error) {
	dbUri := os.Getenv("DATABASE_URL")

	if dbUri == "" {
		dbUri = "host=localhost port=5432 user=naseebullahahmadi dbname=fullstack sslmode=disable"
	}

	d, err := gorm.Open("postgres", dbUri)

	if err != nil {
		return nil, err
	}

	if err = d.DB().Ping(); err != nil {
		return nil, err
	}

	return &DB{d}, nil
}


