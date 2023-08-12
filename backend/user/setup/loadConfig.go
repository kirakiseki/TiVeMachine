package setup

import (
	"os"
	"strconv"
)

type config struct {
	DSN       string
	Port      int
	JWTSecret string
}

var Config config

func LoadConfig() {
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		Inst.Logger.Fatal().Msg("environment variable DSN not set")
	}

	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		Inst.Logger.Warn().Msg("environment variable PORT not set, using default 8080")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		Inst.Logger.Warn().Msg("environment variable PORT not set correctly, using default 8080")
		port = 8080
	}

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		Inst.Logger.Warn().Msg("environment variable JWT_SECRET not set, using default TiVeMachineSecret")
		jwtSecret = "TiVeMachineSecret"
	}

	Config = config{
		DSN:       dsn,
		Port:      port,
		JWTSecret: jwtSecret,
	}
}
