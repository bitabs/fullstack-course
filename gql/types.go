package gql

import . "github.com/graphql-go/graphql"

var User = NewObject(
	ObjectConfig{
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
	},
)
