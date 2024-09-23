package router

import (
	handler "go-api/handlers"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.ListOpeningsHandler)

		v1.GET("/opening", handler.ShowOpeningsHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}
