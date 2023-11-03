package routes

import (
	"databahn-api/config"
	dataSource "databahn-api/data_sources"
	service "databahn-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupLoadRoutes(router *gin.Engine) {
	loadRoute := router.Group("/load")
	{
		loadRoute.GET("/:directory_name/:template_file_name", loadHandler)
	}
}

func loadHandler(c *gin.Context) {
	directoryName := c.Param("directory_name")
	templateFileName := c.Param("template_file_name")
	countParam := c.DefaultQuery("count", "1")
	count, err := strconv.Atoi(countParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid count value"})
		return
	}

	appConfig := config.NewConfig()
	kafkaDataSource, err := dataSource.NewKafkaDataSource(appConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Kafka data source"})
		return
	}

	loadService := service.NewLoadService(kafkaDataSource, appConfig.MaxLimit)

	if err := loadService.LoadData(directoryName, templateFileName, count); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load and process data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loaded and processed data successfully"})
}
