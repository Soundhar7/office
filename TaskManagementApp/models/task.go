package models

import(
	"github.com/jinzhu/gorm"
)

type Task struct{

	gorm.Model

	Title string `json:"title"`
	Description string `json:"description"`
	AssignedTo string `json:"assignedTo"`
	Status string `json:"status"`
}