package routers

import (
	"music/web-music-microservice/app/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	regionsGroup := router.Group("/region")
	{
		regionsGroup.GET("/:name", controllers.GetTracksByRegion)
	}

	return router
}
