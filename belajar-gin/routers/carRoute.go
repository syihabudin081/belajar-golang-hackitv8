package routers

import (
"belajar-gin/controllers"

"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

router := gin.Default()

router.POST("/cars", controllers.CreateCar)
router.PUT("/car/:carID", controllers.UpdateCar)
router.GET("/cars/:carID",controllers.GetCar)
return router
}