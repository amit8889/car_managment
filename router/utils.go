package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UtilsRouter(r *gin.Engine) {
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Health check pass ✅ ✅ ✅ ",
			"success": true,
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.NoRoute(func(c *gin.Context) {
		// Render the 404 page
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"message": "Not found",
			"success": false,
		})
	})
}
