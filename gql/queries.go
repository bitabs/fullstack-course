package gql

import (
	"fullstack-course/db"
	. "github.com/graphql-go/graphql"
)

type Obj struct {
	Query *Object
}

type Queries struct {
	R *Resolver
}

func Query(db *db.DB) *Obj {
	q := Queries{
		R: &Resolver{db:db},
	}

	root := Obj{
		Query: NewObject(ObjectConfig{
			Name: "Query",
			Fields: Fields{
				"tutorial": q.tutorial(),
				"tutorials": q.tutorials(),
			},
		}),
	}

	return &root
}

func (q *Queries) tutorial() *Field {
	field := &Field{
		Type		: Tutorial,
		Description	: "Get Tutorial By ID",
		Args		: FieldConfigArgument{
			"id" 	: &ArgumentConfig{
				Type: Int,
			},
		},
		Resolve: q.R.getTutorial,
	}

	return field
}

func (q *Queries) tutorials() *Field {
	field := &Field{
		Type		: NewList(Tutorial),
		Description	: "Get Tutorial List",
		Resolve		: q.R.getAllTutorials,
	}

	return field
}
