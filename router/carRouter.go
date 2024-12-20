package router

import (
	"database/sql"

	carHandler "github.com/amit8889/car_managemnt/handler/car"
	carService "github.com/amit8889/car_managemnt/service/car"
	store "github.com/amit8889/car_managemnt/store/car"
	"github.com/gin-gonic/gin"
)

func CarRouter(db *sql.DB, r *gin.Engine) {
	carHandler := carInit(db)
	router := r.Group("/car")
	router.GET("/getCarById/:carId", carHandler.GetCarById)
	router.GET("/getAllCars", carHandler.GetAllCars)
	router.POST("/addCar", carHandler.CreateCar)
	//router.PUT("/updateCar/:carId", carHandler.UpdateCar)
	router.DELETE("/deleteCar/:carId", carHandler.DeleteCarById)
	/*
	 cars ID delete
	 cars ID put
	*/
}

func carInit(db *sql.DB) *carHandler.CarHandler {
	carStore := store.New(db)
	carServices := carService.NewCarService(carStore)
	carHandler := carHandler.NewCarHandler(carServices)
	return carHandler
}
