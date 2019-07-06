package gql

import . "github.com/graphql-go/graphql"

var User = NewObject(ObjectConfig{
	Name: "User",
	Fields: Fields{
		"id": &Field{
			Type: Int,
		},
		"name": &Field{
			Type: String,
		},
		"age": &Field{
			Type: Int,
		},
		"profession": &Field{
			Type: String,
		},
		"friendly": &Field{
			Type: Boolean,
		},
	},
})

var Author = NewObject(ObjectConfig{
	Name: "Author",
	Fields: Fields{
		"Name": &Field{
			Type: String,
		},
		"Tutorials": &Field{
			Type: NewList(Int),
		},
	},
})

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
		"author": &Field{
			Type: Author,
		},
		"comments": &Field{
			Type: NewList(Comment),
		},
	},
})
