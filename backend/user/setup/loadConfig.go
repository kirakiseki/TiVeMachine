package setup

import (
	"os"
	"strconv"
)

type config struct {
	DSN           string
	Port          int
	JWTSecret     string
	PasswordSalt  string
	MinioEndpoint string
	MinioAK       string
	MinioSK       string
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

	passwordSalt, ok := os.LookupEnv("PASSWORD_SALT")
	if !ok {
		Inst.Logger.Warn().Msg("environment variable PASSWORD_SALT not set, using default TiVeMachineSecret")
		passwordSalt = "TiVeMachineSecret"
	}

	minioEndpoint, ok := os.LookupEnv("MINIO_ENDPOINT")
	if !ok {
		Inst.Logger.Warn().Msg("environment variable MINIO_ENDPOINT not set, using default minio:9000")
		minioEndpoint = "minio:9000"
	}

	minioAK, ok := os.LookupEnv("MINIO_AK")
	if !ok {
		Inst.Logger.Warn().Msg("environment variable MINIO_AK not set, using default tivemachineak")
		minioAK = "tivemachineak"
	}

	minioSK, ok := os.LookupEnv("MINIO_SK")
	if !ok {
		Inst.Logger.Warn().Msg("environment variable MINIO_SK not set, using default tivemachinesk")
		minioSK = "tivemachinesk"
	}

	Config = config{
		DSN:           dsn,
		Port:          port,
		JWTSecret:     jwtSecret,
		PasswordSalt:  passwordSalt,
		MinioEndpoint: minioEndpoint,
		MinioAK:       minioAK,
		MinioSK:       minioSK,
	}
}
