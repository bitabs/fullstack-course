package gql

import . "github.com/graphql-go/graphql"

var Comment = NewObject(ObjectConfig{
	Name: "Comment",
	Fields: Fields{
		"body": &Field{
			Type: String,
		},
	},
})

var Tutorial = NewObject(ObjectConfig{
	Name: "Tutorial",
	Fields: Fields{
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
