package setup

import "github.com/gin-gonic/gin"

func Gin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	GinEngine := gin.New()
	gin.Default()
	GinEngine.Use(GinLogger()).Use(gin.Recovery())
	Inst.Logger.Info().Msg("Initialized gin")
	return GinEngine
}
