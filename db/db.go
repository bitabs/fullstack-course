package db

import (
	"fmt"
	"fullstack-course/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

type DB struct {
	*gorm.DB
}

func Connect() (*DB, error) {
	dbUri := os.Getenv("DATABASE_URL")

	if dbUri == "" {
		dbUri = "host=localhost port=5432 user=naseebullahahmadi dbname=fullstack sslmode=disable"
	}

	db, err := gorm.Open("postgres", dbUri)

	if err != nil {
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func conStr(isProd bool, config *conf.Config) string {
	// get the database config variables
	c := config.Database

	// local and remote has different config.
	if isProd {
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s", c.Host, c.Port, c.User, c.Name, c.Pass,
		)
	} else {
		return fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Name,
		)
	}
}

func (db *DB) initModel() {
	println("Inside InitModel")
}


