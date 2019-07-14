package db

import (
	"fmt"
	"fullstack-course/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	*gorm.DB
}

func Connect(isProd bool) (*DB, error) {
	config, err := conf.GetConfig(isProd)

	if err != nil {
		log.Fatal("Error::", err)
	}

	db, err := gorm.Open("postgres", conStr(isProd, config))

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


