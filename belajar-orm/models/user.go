package models

import "time"

type User struct {
	ID uint `gorm:"primary_key"`
	Email string `gorm:"type:varchar(100);unique_index"`
	Products []Product
	CreatedAt time.Time
	updateAt time.Time 
}