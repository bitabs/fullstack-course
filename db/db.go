package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

type DB struct {
	*gorm.DB
}

/**
initiates database connection. If app runs on production, get the heroku specific config var
'DATABASE_URL'. Otherwise run a local db.
**/
func (db *DB) Connect() (*DB, error) {

	// on runtime, heroku will set DATABASE_URL with heroku db variables
	dbUri := os.Getenv("DATABASE_URL")

	// run local DB if DATABASE_URL is not set (i.e. if the app is running on dev env
	if dbUri == "" {
		dbUri = "host=localhost port=5432 user=naseebullahahmadi dbname=fullstack sslmode=disable"
	}

	// inits new DB connection
	d, err := gorm.Open("postgres", dbUri)

	// if connection fails
	if err != nil {
		return nil, err
	}

	// Ping verifies a connection to the database is still alive, establishing a connection if
	// necessary.
	if err = d.DB().Ping(); err != nil {
		return nil, err
	}

	// populate our DB struct and return it
	return &DB{d}, nil
}


