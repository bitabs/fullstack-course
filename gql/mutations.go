package gql

import (
	"fullstack-course/models"
	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type: Tutorial,
			Description: "Create new tutorial",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				tut := models.Tutorial{
					Title: p.Args["title"].(string),
				}
				Add(tut)
				return Tuts, nil
			},
		},
	},
})
