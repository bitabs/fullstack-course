package gql

import (
	"fullstack-course/db"
	. "github.com/graphql-go/graphql"
)

type MutationResolver struct {
	R Resolver
}

func Mutation(db *db.DB) *Object {
	res := MutationResolver{
		R: Resolver{db:db},
	}

	root := NewObject(ObjectConfig{
		Name: "Mutation",
		Fields: Fields{
			"createTutorial": res.createTutorial(),
		},
	})

	return root
}

func (mr * MutationResolver) createTutorial() *Field {
	field := &Field{
		Type		: Tutorial,
		Description	: "Create new tutorial",
		Args		: FieldConfigArgument{
			"title"	: &ArgumentConfig{
				Type: NewNonNull(String),
			},
		},
		Resolve: mr.R.createTutorial,
	}

	return field
}

//var Mutation = graphql.NewObject(graphql.ObjectConfig{
//	Name: "Mutation",
//	Fields: graphql.Fields{
//		"create": &graphql.Field{
//			Type: Tutorial,
//			Description: "Create new tutorial",
//			Args: graphql.FieldConfigArgument{
//				"title": &graphql.ArgumentConfig{
//					Type: graphql.NewNonNull(graphql.String),
//				},
//			},
//			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
//				tut := models.Tutorial{
//					Title: p.Args["title"].(string),
//				}
//				return tut, nil
//			},
//		},
//	},
//})
