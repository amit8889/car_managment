package engine

import (
	"github.com/amit8889/car_managemnt/service"
	"github.com/gin-gonic/gin"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(carService service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{carService}
}

func (ch *EngineHandler) GetEngineById(c *gin.Context) {
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
