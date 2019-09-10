package main

import (
	"flag"
	. "fullstack-course/db"
	"fullstack-course/mock"
	"fullstack-course/server"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func init()  {
	var (
		env = flag.String("env", "prod", "application environment")
		database = DB{}
		db, err  = database.Connect()
	)

	gin.SetMode(gin.ReleaseMode)

	flag.Parse()

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	m := mock.Mock{DB: db}

	m.InitModel()

	s := server.Server{
		DB			: db,
		Router		: gin.Default(),
		Environment	: *env,
		PORT		: getPort(),
	}

	s.Server()
}

func main() { }

/**
  * func to fetch available port
  */
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3000"
	}
	return ":" + port
}


