package api

import (
	"fullstack-course/db"
	"fullstack-course/gql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
)

func InitAPI(db *db.DB, r *gin.Engine, env string)  {
	// extract queries
	rootQuery := gql.Query(db).Query
	rootMutation := gql.Mutation(db)

	// our initial single-schema. Queries will hold all of our queries
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		Mutation: rootMutation,
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

	if env == "dev" {
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		r.Use(cors.New(config))
	}

	// let us expose the graphql playground under /graphql endpoint
	r.Group("/graphql")
	{
		// middleware
		r.Use(func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		})
	}
}
