package router

import (
	"go-jwt/controllers"
	"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register",controllers.UserRegister)
		userRouter.POST("/login",controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/create",controllers.CreateProduct)
		productRouter.PUT("/update/:productId",middlewares.ProductAuthorization(),controllers.UpdateProduct)
	}

	return r
}