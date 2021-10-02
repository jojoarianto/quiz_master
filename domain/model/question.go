package model

import (
	"github.com/jinzhu/gorm"
)

// Question is data structure for question entity
type Question struct {
	gorm.Model
	Number   int    `validate:"required" gorm:"type:integer;unique_index" json:"Number"`
	Question string `validate:"required" json:"Question"`
}
