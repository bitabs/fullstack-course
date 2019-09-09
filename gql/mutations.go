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
			"registerUser": res.registerUser(),
		},
	})

	// return our mutation root, so we can register it as part
	// of our single source schema
	return root
}

func (mr *MutationResolver) registerUser() *Field{
	field := &Field{
		Type: User,
		Description: "Register new user",
		Args: FieldConfigArgument{
			"firstName": &ArgumentConfig{
				Type: String,
			},
			"lastName": &ArgumentConfig{
				Type: String,
			},
			"emailAddress": &ArgumentConfig{
				Type: String,
			},
			"phoneNumber": &ArgumentConfig{
				Type: String,
			},
			"username": &ArgumentConfig{
				Type: String,
			},
			"password": &ArgumentConfig{
				Type: String,
			},
		},
		Resolve: mr.R.registerUser,
	}

	return field
}

/**
mutation func that creates a tut.
@Args: Title
**/
func (mr *MutationResolver) createTutorial() *Field {
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
