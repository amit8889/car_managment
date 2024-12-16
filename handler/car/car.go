package car

import (
	"github.com/amit8889/car_managemnt/service"
	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(carService service.CarServiceInterface) *CarHandler {
	return &CarHandler{carService}
}

func (ch *CarHandler) GetCarById(c *gin.Context) {
	// id := r.URL.Query().Get("id")
	// ctx = r.Context()
	// if id == "" {
	// 	http.Error(w, "id is required", http.StatusBadRequest)
	// 	return
	// }
	// car, err := ch.service.GetCarById(id)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// json.NewEncoder(w).Encode(car)

}
