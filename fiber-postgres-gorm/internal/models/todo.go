package models

import "time"

type Todo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
