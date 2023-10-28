package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HomeRouter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ts": time.Now().Unix(),
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health-check": "success",
	})
}
