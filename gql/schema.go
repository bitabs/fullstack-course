package gql

import (
	"fullstack-course/db"
	"github.com/graphql-go/graphql"
)

func TutFields(db *db.DB) graphql.Fields {
	resolver := Resolver{db: db}

	return graphql.Fields{
		"tutorial": &graphql.Field{
			Type		: Tutorial,
			Description	: "Get Tutorial By ID",
			Args		: graphql.FieldConfigArgument{
				"id" : &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.TutorialResolver,
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(Tutorial),
			Description: "Get Tutorial List",
			Resolve: resolver.TutorialsResolver,
		},
	}
}
