package gql

import (
	"fullstack-course/db"
	. "github.com/graphql-go/graphql"
)

type Root struct {
	Query *Object
}

func NewRoot(db *db.DB) *Root {

	resolver := Resolver{db: db}

	root := Root{
		Query: NewObject(
			ObjectConfig{
				Name: "Query",
				Fields: Fields{
					"users": &Field{
						// Slice of User type which can be found in types.go
						Type: NewList(User),
						Args: FieldConfigArgument{
							"name": &ArgumentConfig{
								Type: String,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}

	return &root
}
