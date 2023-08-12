package setup

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GOrm() *gorm.DB {
	DB, err := gorm.Open(mysql.Open(Config.DSN), &gorm.Config{})
	if err != nil {
		Inst.Logger.Fatal().Err(err).Msg("failed to connect database")
	}
	Inst.Logger.Info().Msg("Initialized gorm")
	return DB
}
