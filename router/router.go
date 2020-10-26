package router

import (
	"psys/handler"

	"github.com/gin-gonic/gin"
)

// SetupRoutes - setup router api
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	greenpoint := api.Group("/greenpoint")
	greenpoint.GET("/", handler.GetRedPoint)
	greenpoint.POST("/", handler.AddRedPoint)
}
