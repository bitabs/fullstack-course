package mock

import (
	"fullstack-course/db"
	"fullstack-course/models"
)

type Mock struct { DB *db.DB }

func (mock *Mock) InitModel() {

	// (BEGIN) Below is the code used for dev purposes
	// Two (structs) tables, tutorial & comment
	// TODO: tut, cmt := models.Tutorial{}, models.Comment{}

	// dev related, drop the table if it exists
	// TODO: uncomment to drop tables for dev env mock.DB.DropTableIfExists(&tut, &cmt)

	// Create the table
	// TODO: uncomment to create fresh tables mock.DB.CreateTable(&tut, &cmt)

	// some mock data
	// TODO: Uncomment below for dev env
	//data := models.Tutorial{
	//	Title: "First Tut",
	//	Comments: []models.Comment{
	//		{Body: "First Comment"},
	//		{Body: "Second Comment"},
	//		{Body: "Third Comment"},
	//	},
	//}

	// Insert the data to psql
	//mock.DB.Create(&data)
	// (END)

	var users []models.User

	mock.DB.AutoMigrate(users)
}
