package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"program/controller"
	"program/routers/middleware"
)

func RegisterRoutes(e *gin.Engine) {
	e.Use(middleware.CORS())
	e.Use(cors.Default())

	apiRoutes := e.Group("/api")

	programRoutes := apiRoutes.Group("/program")

	programRoutes.GET("/list", controller.List)
	programRoutes.GET("/channelList", controller.ChannelList)
	programRoutes.GET("/info", controller.Info)
	programRoutes.GET("/channelInfo", controller.ChannelInfo)
	programRoutes.GET("/listByPeriod", controller.ListByPeriod)
	programRoutes.GET("/listByCategory", controller.ListByCategory)
	programRoutes.GET("/listByChannel", controller.ListByChannel)
	programRoutes.GET("/search", controller.Search)
	programRoutes.POST("/addProgram", controller.Add)

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "status": "fail", "msg": "Not found"})
	})
}
