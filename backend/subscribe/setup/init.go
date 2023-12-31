package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Instance struct {
	Logger    *zerolog.Logger
	GinEngine *gin.Engine
	DB        *gorm.DB
}

var Inst Instance

func InitDependencies() {
	Inst.GinEngine = Gin()
	Inst.DB = GOrm()
}

func InitLogger() {
	Inst.Logger = Zerolog()
}
