// Router for the app
package routers

import (
	"github.com/mira-miracoli/brewess2/handlers"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.ErrorLogger())
	resources := router.Group("/resources")
//	resources.DELETE("/:resource/:id", handlers.DeleteResourceByID)
	resources.GET("/hop", handlers.GetAllHops)
//	resources.GET("/:resource/:id", handlers.DisplayResourceByID)
	resources.POST("/hop", handlers.CreateHop)
	return router
}
