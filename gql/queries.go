package gql

import (
	"fullstack-course/db"
	. "github.com/graphql-go/graphql"
)

type Root struct {
	Query *Object
}

func InitRoot(db *db.DB) *Root {
	root := Root{
		Query: NewObject(
			ObjectConfig{
				Name: "Query",
				Fields: TutFields(db),
			},
		),
	}

	return &root
}
