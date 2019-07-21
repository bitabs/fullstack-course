package gql

import (
	"fullstack-course/db"
	"fullstack-course/models"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db	*db.DB
}

func (r *Resolver) getTutorial(p graphql.ResolveParams) (interface{}, error) {
	tut := models.Tutorial{}

	// SELECT * FROM tutorial WHERE ID = p.Args["id"]
	r.db.First(&tut, p.Args["id"].(int))

	// Create a relationship between tutorial and comments
	r.db.Model(&tut).Related(&tut.Comments)

	return tut, nil
}

func (r *Resolver) getAllTutorials(p graphql.ResolveParams) (interface{}, error) {
	var tuts []models.Tutorial

	r.db.Find(&tuts)

	for i := range tuts {
		r.db.Model(&tuts[i]).Related(&tuts[i].Comments)
	}

	return tuts, nil
}

func (r *Resolver) createTutorial(p graphql.ResolveParams) (interface{}, error) {
	tuts := models.Tutorial{
		Title: p.Args["title"].(string),
	}

	r.db.Create(&tuts)

	return tuts, nil
}
