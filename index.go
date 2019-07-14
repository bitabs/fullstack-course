package main

import (
	. "fullstack-course/db"
	"fullstack-course/gql"
	"fullstack-course/models"
	"github.com/gin-contrib/static"
	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type Router struct {
	*gin.Engine
}

func main() {
	// initialise database connection based on the configs
	db := initDBConnection()

	// once we're finished with db. Request the db to close
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// we'll use gin for routing
	router := Router{gin.Default()}

	// initialise Database model, i.e. tables
	initDbModel(db)

	// initialise graphql api
	initGraphQLApi(db, &router)

	// serve static frontend pages
	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	// get a port
	Port := GetPort()

	// log which port we're connected to
	log.Println("Now server is running on port" + Port)

	// initiate the app on available port
	if err := router.Run(Port); err != nil {
		log.Fatal(err)
	}
}

/**
  * func to initialise database model, i.e. create tables
  * within the database.
  */
func initDbModel(db *DB) {
	// Two (structs) tables, tutorial & comment
	tut, cmt := models.Tutorial{}, models.Comment{}

	// dev related, drop the table if it exists
	db.DropTableIfExists(&tut, &cmt)

	// Create the table
	db.CreateTable(&tut, &cmt)

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
	db.Create(&data)
}

/**
  * initialise Database (PostgreSQL) connection.
  */
func initDBConnection() *DB {
	// will begin the connection process
	db, err := Connect()

	// in the event of any failure during the connection, throw the error
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// once connection is established, return the pointer to the database
	return db
}

/**
  * initialise graphql api
  */
func initGraphQLApi(db *DB, router *Router) {
	// first construct a root for our gql schema
	rootQry := gql.InitRoot(db)

	// construct the schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQry.Query,
	})

	// throw the error if it fails
	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	// curated list of gql query
	r := gql.ExecuteQuery(`
		{
			tutorial(id: 1) {
				title
				comments {
					body
				}
			}
			list {
				title
				comments {
					body
				}
			}
		}
	`, schema)

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, r)
		})
	}
}

/**
  * func to fetch available port
  */
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3000"
	}
	return ":" + port
}


