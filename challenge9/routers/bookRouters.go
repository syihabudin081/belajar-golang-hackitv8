package routers

import (
	
	"challenge9/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.GET("/books/", controllers.GetAllBooks)
	router.GET("/books/:bookID", controllers.GetBookByID)
	router.DELETE("/book/:bookID",controllers.DeleteBook)
	return router
}
