package router

import (
	"database/sql"

	engineHandler "github.com/amit8889/car_managemnt/handler/engine"
	engineService "github.com/amit8889/car_managemnt/service/engine"
	"github.com/amit8889/car_managemnt/store/engine"
	"github.com/gin-gonic/gin"
)

func EngineRouter(db *sql.DB, r *gin.Engine) {
	engineHandler := engineInit(db)

	router := r.Group("/engine")
	router.GET("/getEngineById", engineHandler.GetEngineById)

	/*
	 car by id GET
	 cars GET
	 cars POST
	 cars ID delete
	 cars ID put
	*/

}

func engineInit(db *sql.DB) *engineHandler.EngineHandler {

	engineStore := engine.New(db)
	engineServices := engineService.NewEngineService(engineStore)
	engineHandlers := engineHandler.NewEngineHandler(engineServices)
	return engineHandlers
}
