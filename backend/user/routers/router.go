package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user/controller"
	"user/routers/middleware"
)

func RegisterRoutes(e *gin.Engine) {
	e.Use(middleware.AuthInceptor())

	apiRoutes := e.Group("/api")

	userRoutes := apiRoutes.Group("/user")

	userRoutes.GET("/info/:userID", controller.Info)

	userRoutes.POST("/login", controller.Login)
	userRoutes.POST("/register", controller.Register)
	userRoutes.POST("/changeAvatar", controller.ChangeAvatar)

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})
}
