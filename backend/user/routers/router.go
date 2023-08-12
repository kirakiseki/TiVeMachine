package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user/controller"
)

func RegisterRoutes(e *gin.Engine) {
	apiRoutes := e.Group("/api")

	userRoutes := apiRoutes.Group("/user")

	userRoutes.POST("/login", controller.Login)
	userRoutes.POST("/register", controller.Register)
	userRoutes.GET("/info", controller.Info)

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})
}
