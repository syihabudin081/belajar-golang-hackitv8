package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})

}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var updatedCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			CarDatas[i] = updatedCar
			CarDatas[i].CarID = carID
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "data not found",
			"error_message": fmt.Sprintf("car with id %v not found",carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : fmt.Sprintf("car with id %v has succesfully updated",carID),
	})
}

func GetCar(ctx *gin.Context){
carID := ctx.Param("carID")
condition := false
var carData Car

for i,car := range CarDatas{
	if carID == car.CarID {
		condition = true
		carData = CarDatas[i]
		break
	}
}

if !condition {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error_status" : "Data Not Found",
		"error_message" : fmt.Sprintf("car with id %v not found", carID),
	})
	return
}

ctx.JSON(http.StatusOK,gin.H{
	"car" : carData,
})

}