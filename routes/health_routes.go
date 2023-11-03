package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupHealthRoutes(router *gin.Engine) {
	router.GET("/ping", healthHandler)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
