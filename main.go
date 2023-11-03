package main

import (
	"databahn-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupLoadRoutes(r)
	routes.SetupHealthRoutes(r)

	r.Run(":8080")
}
