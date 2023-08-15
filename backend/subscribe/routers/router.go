package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"subscribe/controller"
	"subscribe/routers/middleware"
)

func RegisterRoutes(e *gin.Engine) {
	e.Use(middleware.AuthInterceptor())
	e.Use(middleware.CORS())
	e.Use(cors.Default())

	apiRoutes := e.Group("/api")

	subscribeRoutes := apiRoutes.Group("/subscribe")

	subscribeRoutes.GET("/subscriptionList", controller.SubscriptionList)
	subscribeRoutes.POST("/subscribe", controller.Subscribe)
	subscribeRoutes.POST("/unsubscribe", controller.Unsubscribe)

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "status": "fail", "msg": "Not found"})
	})
}
