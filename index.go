package main

import (
	"flag"
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

	boolPtr := flag.Bool("prod", false, "for production")

	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	router := Router{gin.Default()}

	// initialise database connection based on the configs
	db := initDBConnection(*boolPtr)

	// once we're finished with db. Request the db to close
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Two tables, tutorial & comment
	tut, cmt := models.Tutorial{}, models.Comment{}

	// dev purposes, drop the table if it exists
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

	initGraphQLApi(db, &router)

	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	// get a port
	Port := GetPort()

	// log which port we're connected to
	log.Println("Now server is running on port" + Port)

	if err := router.Run(Port); err != nil {
		log.Fatal(err)
	}
}

// Initialises Database (PostgreSQL) connection.
func initDBConnection(isProd bool) *DB {
	// will begin the connection process
	db, err := Connect(isProd)

	// in the event of any failure during the connection,
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// once connection is established, return the pointer to the database
	return db
}

func initGraphQLApi(db *DB, router *Router)  {
	rootQry := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: gql.TutFields(db),
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQry),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	r := gql.ExecuteQuery(`
		{
			tutorial(id: 1) {
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


func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3000"
	}
	return ":" + port
}


