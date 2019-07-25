package gql

import (
	"fullstack-course/db"
	. "github.com/graphql-go/graphql"
)

// our struct that will hold resolver
type MutationResolver struct { R Resolver }

/**
GraphQL specific mutation function will link mutation functions
**/
func Mutation(db *db.DB) *Object {

	// populate the struct with the resolver and passing the
	// db connection,
	res := MutationResolver{ R: Resolver{db:db} }

	// the root of mutation object
	root := NewObject(ObjectConfig{
		Name: "Mutation",
		Fields: Fields{
			// [0..*] of mutations
			"createTutorial": res.createTutorial(),
		},
	})

	// return our mutation root, so we can register it as part
	// of our single source schema
	return root
}

/**
mutation func that creates a tut.
@Args: Title
**/
func (mr * MutationResolver) createTutorial() *Field {
	field := &Field{
		Type:        Tutorial,
		Description: "Create new tutorial",
		Args: FieldConfigArgument{
			"title": &ArgumentConfig{
				Type: NewNonNull(String),
			},
		},
		// calls the func in resolver that communicates to the db,
		// and creates the tut
		Resolve: mr.R.createTutorial,
	}

	return field
}
