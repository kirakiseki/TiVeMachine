package main

import (
	"fmt"
	"program/routers"
	"program/setup"
)

func main() {
	setup.InitLogger()
	setup.LoadConfig()
	setup.InitDependencies()

	routers.RegisterRoutes(setup.Inst.GinEngine)

	err := setup.Inst.GinEngine.Run(fmt.Sprintf(":%d", setup.Config.Port))

	if err != nil {
		setup.Inst.Logger.Fatal().Err(err).Msg("failed to start server")
	}
}
