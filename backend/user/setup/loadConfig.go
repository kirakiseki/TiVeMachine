package setup

import (
	"os"
	"strconv"
)

type config struct {
	DSN  string
	Port int
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

	Config = config{
		DSN:  dsn,
		Port: port,
	}
}
