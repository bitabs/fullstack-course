package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	TutorialID	uint
	Body 		string
}
