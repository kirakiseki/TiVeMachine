package setup

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Instance struct {
	Logger       *zerolog.Logger
	GinEngine    *gin.Engine
	DB           *gorm.DB
	MinioClient  *minio.Client
	MinioContext *context.Context
}

var Inst Instance

func InitDependencies() {
	Inst.GinEngine = Gin()
	Inst.DB = GOrm()
	Inst.MinioClient, Inst.MinioContext = Minio()
}

func InitLogger() {
	Inst.Logger = Zerolog()
}
