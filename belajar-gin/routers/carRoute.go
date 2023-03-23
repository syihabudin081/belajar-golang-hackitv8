package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.PUT("/car/:carID", controllers.UpdateCar)
	router.GET("/cars/", controllers.GetAllCars)
	router.GET("/cars/:carID", controllers.GetCarByID)
	router.DELETE("/car/:carID",controllers.DeleteCar)
	return router
}
