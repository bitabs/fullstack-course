package gql

import (
	"fullstack-course/models"
)

var Tuts []models.Tutorial

func Populate() []models.Tutorial {
	author := &models.Author{
		Name: "Naseebullah Ahmadi",
		Tutorials: []int{1, 2},
	}

	tutorial1 := models.Tutorial{
		Title: "Go GraphQL Tutorial",
		Author: *author,
		Comments: []models.Comment{
			{Body: "First Comment"},
		},
	}

	tutorial2 := models.Tutorial{
		Title: "Go GraphQL Tutorial - Part 2",
		Author: *author,
		Comments: []models.Comment{
			{Body: "Second Comment"},
		},
	}

	Tuts = append(Tuts, tutorial1)
	Tuts = append(Tuts, tutorial2)

	return Tuts
}

func Add(tutorial models.Tutorial) {
	Tuts = append(Tuts, tutorial)
}
