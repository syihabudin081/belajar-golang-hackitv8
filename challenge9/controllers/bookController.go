package controllers

import (
	"challenge9/database"
	"challenge9/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.GetDB().Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book created successfully",
		"book":    book,
	})
}

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	if err := database.GetDB().Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Books retrieved successfully",
		"books":   books,
	})
}

func UpdateBook(c *gin.Context) {
	bookID := c.Param("bookID")
	var book models.Book
	if err := database.GetDB().First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedbook models.Book
	if err := c.BindJSON(&updatedbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.GetDB().Model(&book).Updates(updatedbook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
		"book":    book,
		
	})

}


func DeleteBook(c *gin.Context) {
	bookID := c.Param("bookID")
	var book models.Book
	if err := database.GetDB().First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := database.GetDB().Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}

func GetBookByID(c *gin.Context) {
	bookID := c.Param("bookID")
 book := models.Book{}

	if err := database.GetDB().First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book retrieved successfully",
		"book":    book,
	})

}