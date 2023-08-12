package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user/controller"
)

func RegisterRoutes(e *gin.Engine) {
	e.Use(AuthInceptor())

	apiRoutes := e.Group("/api")

	userRoutes := apiRoutes.Group("/user")
	userRoutes.POST("/login", controller.Login)
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/changeAvatar", controller.ChangeAvatar)
	userRoutes.GET("/info/:userID", controller.Info)

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})
}
