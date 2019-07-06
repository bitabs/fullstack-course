package main

import (
	. "fullstack-course/db"
	"fullstack-course/models"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"os"
)

func initAPI() ( *chi.Mux, *DB )  {
	r := chi.NewRouter()

	if e := godotenv.Load(); e != nil {
		log.Fatalf("Failed to load env file: %s", e)
	}

	Host 		:= os.Getenv("Host")
	Database 	:= os.Getenv("Database")
	User		:= os.Getenv("User")
	Port		:= os.Getenv("Port")
	Password	:= os.Getenv("Password")

	db, err := Connect(ConnectionStr(Host, Port, User, Database, Password))

	if err != nil {
		log.Fatal( err)
	}

	//db.AutoMigrate(&models.Author{}, &models.Comment{}, &models.Tutorial{})

	//rootQuery := gql.NewRoot(db)
	//
	//sc, err := graphql.NewSchema(
	//	graphql.SchemaConfig{Query: rootQuery.Query},
	//)
	//
	//if err != nil {
	//	fmt.Println("Error creating schema: ", err)
	//}
	//
	//s := server.Server{
	//	GqlSchema: &sc,
	//}
	//
	//// Add some middleware to our router
	//r.Use(
	//	render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
	//	middleware.Logger,          // log api request calls
	//	middleware.DefaultCompress, // compress results, mostly gzipping assets and json
	//	middleware.StripSlashes,    // match paths with a trailing slash, strip it, and continue routing through the mux
	//	middleware.Recoverer,       // recover from panics without crashing server
	//)

	r.Post("/graphql", nil)
	//r.Post("/graphql", s.GraphQL())
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

func main() {
	// get a port
	Port := GetPort()

	_, db := initAPI()

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//db.AutoMigrate(&models.Tutorial{})

	//db.Create(&models.Tutorial{
	//	Title: "My First Tut",
	//	Comments: []string{
	//		"Hello World",
	//	},
	//})

	//db.Model(&models.Tutorial{}).Related(&models.Comment{})


	//db.CreateTable(&models.Tutorial{})

	//rootQuery := graphql.ObjectConfig{
	//	Name: "RootQuery",
	//	Fields: gql.TutFields(db),
	//}
	//
	//schemaConfig := graphql.SchemaConfig{
	//	Query: graphql.NewObject(rootQuery),
	//}
	//
	//schema, err := graphql.NewSchema(schemaConfig)
	//
	//if err != nil {
	//	log.Fatalf("Failed to create new schema, error: %v", err)
	//}
	//
	//query := `
	//	{
	//		tutorial(id: 1) {
	//			title,
	//			author
	//		}
	//	}
	//`
	//
	//params := graphql.Params{
	//	Schema: schema,
	//	RequestString: query,
	//}
	//
	//r := graphql.Do(params)
	//
	//if len(r.Errors) > 0 {
	//	log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	//}
	//
	//rJSON, _ := json.Marshal(r)
	//
	//fmt.Printf("%s \n: ", rJSON)


	// serve static Frontend React pages
	http.Handle("/", http.FileServer(http.Dir("./build")))

	// log which port we're connected to
	log.Println("Now server is running on port" + Port)

	// serve our BE to the port, and log any error if it fails to connect to port
	if err := http.ListenAndServe("127.0.0.1" + Port, nil); err != nil {
		log.Fatal(err)
	}
}
