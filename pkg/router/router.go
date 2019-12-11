package router

import (
	"LedImageView-WebAPI/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api/v1")
	{
		upload := api.Group("/upload")
		{
			upload.POST("/", controllers.UploadAnimation)
		}
	}

	return router
}