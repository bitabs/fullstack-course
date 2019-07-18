package mock

import (
	"fmt"
	"fullstack-course/db"
	"fullstack-course/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

type Mock struct {
	DB	*db.DB
}

func (mock *Mock) InitModel() {

	_ = gin.Default()

	var wg sync.WaitGroup

	wg.Add(1)

	go func(url string) {

		// Decrement the counter when the go routine completes
		defer wg.Done()

		// Call the function check
		checkUrl(url)

	}("https://randomuser.me/api/")

	// Two (structs) tables, tutorial & comment
	tut, cmt := models.Tutorial{}, models.Comment{}

	// dev related, drop the table if it exists
	mock.DB.DropTableIfExists(&tut, &cmt)

	// Create the table
	mock.DB.CreateTable(&tut, &cmt)

	// some mock data
	data := models.Tutorial{
		Title: "First Tut",
		Comments: []models.Comment{
			{Body: "First Comment"},
			{Body: "Second Comment"},
			{Body: "Third Comment"},
		},
	}

	// Insert the data to psql
	mock.DB.Create(&data)
}

func checkUrl(url string) {

	_, err := http.Get(url)

	if err != nil {

		log.Fatal("Error! URL: " + url + " is down!")

		return
	}

	fmt.Println(url, "is up and running...")
}
