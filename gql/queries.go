package gql

import (
	"fullstack-course/db"
	. "github.com/graphql-go/graphql"
)

// struct that points to graphql Object
type Obj struct { Query *Object }

// struct that points to Resolver
type Queries struct { R *Resolver }

/**
func that inits our query blueprint for our schema
**/
func Query(db *db.DB) *Obj {

	// populate our Query struct
	q := Queries{ R: &Resolver{db:db} }

	// graphql query object
	root := Obj{
		Query: NewObject(ObjectConfig{
			Name: "Query",
			Fields: Fields{
				// [0..*] fetch queries
				"tutorial": q.tutorial(), 	// fetch particular query: [0...1]
				"tutorials": q.tutorials(), // fetch all queries: [0..*]
			},
		}),
	}

	return &root
}

/**
query that fetches particular tutorial from the DB
**/
func (q *Queries) tutorial() *Field {

	// graphql field
	field := &Field{
		Type		: Tutorial,
		Description	: "Get Tutorial By ID",
		Args		: FieldConfigArgument{
			// id used to search for particular tutorial
			"id" 	: &ArgumentConfig{
				Type: Int,
			},
		},
		// call getTutorial func from resolver that
		// requests the tut from DB
		Resolve: q.R.getTutorial,
	}

	return field
}

/**
query that fetches all tutorials from the DB
**/
func (q *Queries) tutorials() *Field {

	// graphql field
	field := &Field{
		Type		: NewList(Tutorial),
		Description	: "Get Tutorial List",
		// return all tutorials
		Resolve		: q.R.getAllTutorials,
	}

	return field
}
