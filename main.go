package main

import (
	"psys/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	router.SetupRoutes(gin)
	gin.Run(":4000")
}
