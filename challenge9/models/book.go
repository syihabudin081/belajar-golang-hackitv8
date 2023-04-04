package models

import (
	"time"
)

type Book struct {
	ID uint `gorm:"primary_key"`
	Name string `gorm:"not null;type:varchar(100)"`
	Author string `gorm:"not null;type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time 
}