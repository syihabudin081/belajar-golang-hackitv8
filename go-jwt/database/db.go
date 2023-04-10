package database

import (
	"fmt"
	"go-jwt/models"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host	 = "localhost"
	user = "postgres"
	password = "123"
	dbPort = "5432"
	dbname = "simple_api"
	db *gorm.DB
	err error
)


func StartDB(){

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",host,user,password,dbname,dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully")
	db.Debug().AutoMigrate(&models.User{},&models.Product{})
}


func GetDB() *gorm.DB {
	return db
}