package gql

import . "github.com/graphql-go/graphql"

// declare the comment blueprint
var Comment = NewObject(ObjectConfig{
	Name: "Comment",
	Fields: Fields{
		// each comment will have body field
		"body": &Field{
			// of type string
			Type: String,
		},
	},
})

// declare Tutorial blueprint
var Tutorial = NewObject(ObjectConfig{
	Name: "Tutorial",
	Fields: Fields{
		// each tutorial will have the following fields
		"id": &Field{
			Type: Int,
		},
		"title": &Field{
			Type: String,
		},
		"comments": &Field{
			Type: NewList(Comment),
		},
	},
})
