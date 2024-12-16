package response

import "github.com/gin-gonic/gin"

func WriteJson(c *gin.Context, status int, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.Status(status)
	c.JSON(status, data)
}
