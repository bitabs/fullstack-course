package models

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	TutorialID	uint
	Name 		string
	Tutorials 	[]int
}
