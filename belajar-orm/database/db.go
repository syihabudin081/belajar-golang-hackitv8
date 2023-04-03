package database

import (
	"belajar-orm/models"
	"fmt"
	
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	host = "localhost"
	user = "postgres"
	port = "5432"
	password = "123"
	dbname = "learning_gorm"
	db *gorm.DB
	err error
)

func StartDB(){
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database",err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})

}

func GetDB() *gorm.DB {
	return db
}