package main

import (
	"databahn-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupHealthRoutes(r)
	routes.SetupLoadRoutes(r)

	r.Run(":8080")
}
