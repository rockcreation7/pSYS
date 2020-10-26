package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRedPoint -
func GetRedPoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "message",
		"nick":    "nick",
	})
}

// AddRedPoint -
func AddRedPoint(c *gin.Context) {
	type redPointInput struct {
		CustomerID string `form:"customerid" json:"customerid" binding:"required"`
		Point      int    `form:"point" json:"point" binding:"required"`
	}
	var form redPointInput

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  "saved",
		"message": "success",
		"data":    form,
	})
}
