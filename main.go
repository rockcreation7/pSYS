package main

import (
	"io"
	"os"
	"psys/router"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	loggerConfig()
	gin := gin.Default()
	router.SetupRoutes(gin)
	gin.Use(static.Serve("/", static.LocalFile("./dist", true)))
	gin.NoRoute(redirectToSPA)
	gin.Run(":4000")
}

func redirectToSPA(c *gin.Context) {
	c.File("./dist/index.html")
}

func loggerConfig() {
	currentTime := time.Now()
	f, _ := os.Create("log/gin" + currentTime.Format("20060102150405") + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
