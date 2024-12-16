package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB, r *gin.Engine) {
	CarRouter(db, r)
	EngineRouter(db, r)
	UtilsRouter(r)
}
