package main

import (
	"fmt"
	. "fullstack-course/db"
	"fullstack-course/gql"
	"fullstack-course/server"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-chi/chi"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
	"os"
)

func main() {
	// get a port
	Port := GetPort()

	r, db := initAPI()

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// serve static Frontend React pages
	http.Handle("/", http.FileServer(http.Dir("./build")))

	// log which port we're connected to
	log.Println("Now server is running on port" + Port)

	// serve our BE to the port, and log any error if it fails to connect to port
	if err := http.ListenAndServe("127.0.0.1" + Port, r); err != nil {
		log.Fatal(err)
	}
}

func initAPI() ( *chi.Mux ,*DB)  {
	r := chi.NewRouter()

	db, err := New(ConnectionStr("localhost", 5432, "naseebullahahmadi", "go_graphql_db"))

	if err != nil {
		log.Fatal( err)
	}

	rootQuery := gql.NewRoot(db)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)

	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{
		GqlSchema: &sc,
	}

	// Add some middleware to our router
	r.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,          // log api request calls
		middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes,    // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,       // recover from panics without crashing server
	)

	//r.HandleFunc("/graphql", s.GraphQL())
	r.Post("/graphql", s.GraphQL())
	return r, db
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3000"
	}
	return ":" + port
}
