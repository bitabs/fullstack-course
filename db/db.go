package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type DB struct {
	*gorm.DB
}

func Connect(url string) (*DB, error) {
	db, err := gorm.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func ConnectionStr(host string, port string, user string, dbName string, password string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbName, password,
	)
}

func (db *DB) initModel() {
	println("Inside InitModel")
}


//func (d *DB) GetUsersByName(name string) []User  {
//	// Prepare query, takes a name argument, protects from sql injection
//	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
//	if err != nil {
//		fmt.Println("GetUserByName Preperation Err: ", err)
//	}
//
//	// Make query with our stmt, passing in name argument
//	rows, err := stmt.Query(name)
//	if err != nil {
//		fmt.Println("GetUserByName Query Err: ", err)
//	}
//
//	// Create User struct for holding each row's data
//	var r User
//	// Create slice of Users for our response
//	users := []User{}
//	// Copy the columns from row into the values pointed at by r (User)
//	for rows.Next() {
//		err = rows.Scan(
//			&r.ID,
//			&r.Name,
//			&r.Age,
//			&r.Profession,
//			&r.Friendly,
//		)
//		if err != nil {
//			fmt.Println("Error scanning rows: ", err)
//		}
//		users = append(users, r)
//	}
//
//	return users
//}


