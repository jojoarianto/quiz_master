package model

import (
	"github.com/jinzhu/gorm"
)

// Answer is data structure for answer entity
type Answer struct {
	gorm.Model
	Answer int `validate:"required" gorm:"type:varchar(100);unique_index" json:"Answer"`
}
