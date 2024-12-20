package car

import (
	"strconv"

	"github.com/amit8889/car_managemnt/models"
	"github.com/amit8889/car_managemnt/service"
	"github.com/amit8889/car_managemnt/utils/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(carService service.CarServiceInterface) *CarHandler {
	return &CarHandler{carService}
}

func (ch *CarHandler) GetCarById(c *gin.Context) {
	carId, err := uuid.Parse(c.Param("carId"))
	var car models.Car
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid car id", "success": false, "value": car})
		return
	}
	// uuuid to string
	carIdStr := carId.String()
	car, err = ch.service.GetCarById(c, carIdStr)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error(), "success": false, "value": car})
		return
	}
	c.JSON(200, gin.H{"message": "car details!!", "success": true, "value": car})

}
func (ch *CarHandler) GetAllCars(c *gin.Context) {
	//USE PAGINATION
	var cars []models.Car
	var total int32
	page, err := strconv.ParseInt(c.DefaultQuery("page", "0"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error(), "success": false, "cars": cars, "count": 0})
		return
	}

	cars, total, err = ch.service.GetAllCars(c, int32(page))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error(), "success": false, "cars": "cars", "count": total})
		return
	}
	c.JSON(200, gin.H{"message": "all cars!!", "success": true, "value": cars, "count": total})

}

func (ch *CarHandler) CreateCar(c *gin.Context) {
	var car models.Car
	//err := validator.ValidateRequest(c, car)
	err := c.BindJSON(&car)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error(), "success": false})
		return
	}
	//gernerte uuid
	car.ID, err = uuid.NewRandom()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error(), "success": false})
		return
	}
	car.Engine.EngineID, err = uuid.NewRandom()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error(), "success": false})
		return
	}

	errors := validator.ValidateStruct(car)
	if errors != nil {
		c.JSON(400, gin.H{"message": "Invalid request!!", "success": false, "errors": errors})
		return
	}

	err = ch.service.CreateCar(c, car)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error(), "success": false})
		return
	}
	c.JSON(200, gin.H{"message": "car created!!", "success": true})
}
func (ch *CarHandler) DeleteCarById(c *gin.Context) {
	carId := c.Param("carId")
	err := uuid.Validate(carId)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid car id!!", "success": false})
	}
	err = ch.service.DeleteCarById(c, uuid.MustParse(carId))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error(), "success": false})
		return
	}
	c.JSON(200, gin.H{"message": "car deleted!!", "success": true})

}
