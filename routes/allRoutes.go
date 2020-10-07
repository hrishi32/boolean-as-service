package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hrishi32/boolean-as-service/controller"
)

// Init function sets all routes to the server.
func Init(server *gin.Engine) {

	server.GET("/:id", controller.GetHandler)

	server.POST("/", controller.PostHandler)

	server.PATCH("/:id", controller.PatchHandler)

	server.DELETE("/:id", controller.DeleteHandler)

	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

}
