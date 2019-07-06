package gql

import (
	"fullstack-course/db"
	"fullstack-course/models"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *db.DB
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	//name, ok := p.Args["name"].(string)
	//if ok {
		//users := r.db.GetUsersByName(name)
		//return users, nil
	//}
	return nil, nil
}

func (r *Resolver) TutorialResolver(p graphql.ResolveParams) (interface{}, error) {
	auth := &models.Author{
		Name: "Naseebullah",
		Tutorials: []int{1, 2},
	}

	tut := models.Tutorial{
		Title: "Golang",
		Author: *auth,
		Comments: []models.Comment{
			{Body: "First Comment"},
		},
	}

	r.db.Create(&tut)
	//r.db.AutoMigrate(&tut)
	//r.db.Create(&models.Tutorial{
	//	Title: "Testing",
	//	Author: models.Author{
	//		Name: "Naseebullah",
	//		Tutorials: []int{1, 2},
	//	},
	//})
	//r.db.First(&tut, p.Args["id"].(int))
	return tut, nil
}

func (r *Resolver) TutorialsResolver(p graphql.ResolveParams) (interface{}, error) {
	var tuts []models.Tutorial
	r.db.AutoMigrate(&tuts)
	r.db.First(&tuts)
	return tuts, nil
}
