package gql

import (
	"fullstack-course/db"
	. "github.com/graphql-go/graphql"
)

type Root struct {
	Query *Object
}

func InitRoot(db *db.DB) *Root {
	resolver := Resolver{db: db}

	root := Root{
		Query: NewObject(ObjectConfig{
			Name: "Query",
			Fields: Fields{
				"tutorial": &Field{
					Type		: Tutorial,
					Description	: "Get Tutorial By ID",
					Args		: FieldConfigArgument{
						"id" : &ArgumentConfig{
							Type: Int,
						},
					},
					Resolve: resolver.TutorialResolver,
				},
				"list": &Field{
					Type:        NewList(Tutorial),
					Description: "Get Tutorial List",
					Resolve: resolver.TutorialsResolver,
				},
			},
		}),
	}

	return &root
}
