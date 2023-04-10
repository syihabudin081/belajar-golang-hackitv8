package middlewares

import (
	"go-jwt/database"
	"go-jwt/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {

	return func(c *gin.Context) {
		
		db := database.GetDB()
		productId,err := strconv.Atoi(c.Param("productId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message" : "data not found",
			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"error" : "Unauthorized",
				"message" : "You are not authorized to access this resource",
			})
			return
		}

		c.Next()


	}

}