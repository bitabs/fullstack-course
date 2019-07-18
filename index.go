package main

import (
	. "fullstack-course/db"
	"fullstack-course/gql"
	"fullstack-course/mock"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

type Router struct {
	*gin.Engine
}

func main() {
	// declare a variable for our GORM DB model
	database, router := DB{}, Router{gin.Default()}

	// initiate database connection
	db, err := database.Connect()

	// in the event of any failure during the connection, throw the error
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// once we're finished with db. Request the db to close
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	/**
	  * An entry point to initialise database model. i.e,
	  * construct. Note only modify database model at this stage
	 **/
	m := mock.Mock{ DB: db }
	m.InitModel()

	// serve our react frontend
	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	// initialise graphql api
	router.initGraphQLApi(db)

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
  * initialise graphql api
  */
func (r *Router) initGraphQLApi(db *DB) {
	//first construct a root for our gql schema
	rootQry := gql.InitRoot(db)

	// construct the schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQry.Query,
	})

	//throw the error if it fails
	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		Playground: true,
	})

	r.Group("/graphql")
	{
		r.Use(func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
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


