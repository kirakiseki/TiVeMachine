package setup

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Minio() (*minio.Client, *context.Context) {
	MCtx := context.Background()

	MClient, err := minio.New(Config.MinioEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(Config.MinioAK, Config.MinioSK, ""),
	})
	if err != nil {
		Inst.Logger.Fatal().Err(err).Msg("failed to connect to minio")
	}

	Inst.Logger.Info().Msg("Initialized minio")
	return MClient, &MCtx
}
