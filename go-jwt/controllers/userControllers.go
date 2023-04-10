package controllers

import (
	"github.com/gin-gonic/gin"
	"go-jwt/database"
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType != appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User not created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    http.StatusOK,
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
	})

}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType != appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.PassWord
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid email or password",
			"error":   err.Error(),
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.PassWord), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid email or password",
			"error":   err.Error(),
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  token,
	})

}
