package model

import (
	"time"
)

// Question is data structure for question entity
type Question struct {
	ID        uint   `gorm:"primaryKey"`
	Number    int    `validate:"required" gorm:"type:integer;index" json:"Number"`
	Question  string `validate:"required" json:"Question"`
	Answer    string `validate:"required" gorm:"type:string;" json:"answer"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
