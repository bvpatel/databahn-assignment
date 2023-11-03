package routes

import (
	"databahn-api/config"
	"databahn-api/data"
	"databahn-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupLoadRoute(router *gin.Engine) {
	loadRoute := router.Group("/load")
	{
		loadRoute.GET("/:directory_name/:template_file_name", loadTemplate)
	}
}

func loadTemplate(c *gin.Context) {
	directoryName := c.Param("directory_name")
	templateFileName := c.Param("template_file_name")
	countParam := c.DefaultQuery("count", "1")
	count, err := strconv.Atoi(countParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid count value"})
		return
	}

	appConfig := config.NewConfig()
	kafkaDataSource, err := data.NewKafkaDataSource(appConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Kafka data source"})
		return
	}

	loadService := service.NewLoadService(kafkaDataSource, count)

	if err := loadService.LoadData(directoryName, templateContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load and process data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loaded and processed data successfully"})
}
