package gql

import (
	"fullstack-course/db"
	"fullstack-course/models"
	"github.com/graphql-go/graphql"
)

// struct that points to our DB which will have a connection instance
type Resolver struct { db *db.DB }

/**
function that will fetch the tutorial from the database
@params p <graphql.ResolveParams> the params of this query (i.e. 'id')
**/
func (r *Resolver) getTutorial(p graphql.ResolveParams) (interface{}, error) {

	// our tutorial struct that holds the blueprint of what it should contain
	tut := models.Tutorial{}

	// SELECT * FROM tutorial WHERE ID = p.Args["id"]
	r.db.First(&tut, p.Args["id"].(int))

	// Create a relationship between tutorial and comments
	r.db.Model(&tut).Related(&tut.Comments)

	return tut, nil
}

func (r *Resolver) getAllTutorials(p graphql.ResolveParams) (interface{}, error) {

	// declare empty tutorial model
	var tuts []models.Tutorial

	// SELECT * FROM tutorials
	r.db.Find(&tuts)

	// loop through each of the tutorial, and create relationship between itself
	// and the corresponding comments array
	for i := range tuts {
		r.db.Model(&tuts[i]).Related(&tuts[i].Comments)
	}

	return tuts, nil
}

/**
function that will create tutorial in our DB
**/
func (r *Resolver) createTutorial(p graphql.ResolveParams) (interface{}, error) {

	// our tutorial model containing the title
	tuts := models.Tutorial{ Title: p.Args["title"].(string) }

	// INSERT INTO tutorials("title") values( p.Args[ "title" ] );
	r.db.Create(&tuts)

	return tuts, nil
}
