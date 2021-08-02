package models

import "gorm.io/gorm"


type TaskModel struct {
	gorm.Model
	Todo string
	Completed bool
	User string
}





