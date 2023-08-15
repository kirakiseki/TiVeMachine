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
	programRoutes.GET("/scheduleList", controller.ScheduleList)
	programRoutes.POST("/info", controller.Info)
	programRoutes.POST("/channelInfo", controller.ChannelInfo)
	programRoutes.POST("/scheduleInfo", controller.ScheduleInfo)
	programRoutes.POST("/listByPeriod", controller.ListByPeriod)
	programRoutes.POST("/listByCategory", controller.ListByCategory)
	programRoutes.POST("/listByChannel", controller.ListByChannel)
	programRoutes.POST("/search", controller.Search)
	programRoutes.POST("/addProgram", controller.Add)

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "status": "fail", "msg": "Not found"})
	})
}
