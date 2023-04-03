package main

import (
	"belajar-orm/database"
	"belajar-orm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func main() {	

	database.StartDB()
	// createUser("popol@email.com")
	// getUserByID(1)
	// updateUserByID(1, "mega@email.com")
	// createProduct("B", "K",1)
	// getUsersWithProducts()
	deleteProductWithID(1)
}

func createUser(email string) {
	db := database.GetDB()
	user := models.User{
		Email: email,
	}
	err := db.Create(&user).Error

	if err != nil {
		fmt.Println("error creating user", err)
	}
	return
}

func getUserByID(id uint){

	db := database.GetDB()
	user:=  models.User{}
	err := db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			fmt.Println("user not found")
			return
		}
		print("error getting user", err)
	}
	fmt.Println(user)

	db.First(&user, id)
	
}

func updateUserByID(id uint, email string) {
	db := database.GetDB()
	user := models.User{}
	err := db.Model(&user).Where("id = ?", id).Update("email", email).Error

	if err != nil {
		fmt.Println("error updating user", err)
		return
	}
	fmt.Println("Updated User Email", user.Email)
}

func createProduct(name, brand string, userID uint) {
	db := database.GetDB()
	product := models.Product{
		Name: name,
		Brand: brand,
		UserID: userID,
	}
	err := db.Create(&product).Error

	if err != nil {
		fmt.Println("error creating product", err)
	return
	}

	fmt.Println("Created Product", product)
}


func getUsersWithProducts(){
	db := database.GetDB()
	users := []models.User{}
	err := db.Preload("Products").Find(&users).Error
	if err != nil {
		fmt.Println("error getting users", err)
		return
	}
	fmt.Println(users)
}

func deleteProductWithID(id uint){
	db := database.GetDB()
	product := models.Product{}
	err := db.Delete(&product, id).Error
	if err != nil {
		fmt.Println("error deleting product", err)
		return
	}
	fmt.Println("Deleted product")
}