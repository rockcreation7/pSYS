package router

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes setup router api
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	greenpoint := api.Group("/greenpoint")
	greenpoint.GET("/", getGreenPoint)
}

func getGreenPoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "message",
		"nick":    "nick",
	})
}
