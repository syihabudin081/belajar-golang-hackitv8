package controllers

import (
	"fmt"
	"go-jwt/database"
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	useriD := userData["id"].(float64)

	if contentType != appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = uint(useriD)
	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, Product)

}

func UpdateProduct(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType != appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description,}).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, Product)

}


func GetProductByID(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))


	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).First(&Product).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, Product)

}


func GetProducts(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	products := []models.Product{}
	isAdmin, ok := userData["admin"].(bool)
	fmt.Println(userData)
	fmt.Println("isAdmin", isAdmin)

	if !ok {
		isAdmin = false
	}

	if isAdmin {
		err := db.Find(&products).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, products)
		return
	}

	err := db.Where("user_id = ?", userID).Find(&products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}


func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Where("id = ?", productId).Delete(&Product).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Product Deleted Successfully!",
	})

}