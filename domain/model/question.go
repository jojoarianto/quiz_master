package model

import (
	"time"
)

// Question is data structure for question entity
type Question struct {
	ID        uint   `gorm:"primaryKey"`
	Number    int    `validate:"required" gorm:"type:integer;unique_index" json:"Number"`
	Question  string `validate:"required" json:"Question"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
