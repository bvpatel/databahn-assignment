package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupHealthRoutes(router *gin.Engine) {
	router.GET("/health", healthHandler)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
