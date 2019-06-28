package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

type DB struct {
	*sql.DB
}

func New(connectionStr string) (*DB, error) {
	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func ConnectionStr(host string, port int, user string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable", host, strconv.Itoa(port), user, dbName,
	)
}

type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func (d *DB) GetUsersByName(name string) []User  {
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUserByName Preperation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUserByName Query Err: ", err)
	}

	// Create User struct for holding each row's data
	var r User
	// Create slice of Users for our response
	users := []User{}
	// Copy the columns from row into the values pointed at by r (User)
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}

	return users
}


