package utils

import (
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"time"
	"user/setup"
)

func UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64) error {
	updInfo, err := setup.Inst.MinioClient.PutObject(*setup.Inst.MinioContext, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})

	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("failed to upload file")
		return err
	}

	setup.Inst.Logger.Info().Msgf("Successfully uploaded file %s of size %d", updInfo.Key, updInfo.Size)
	return nil
}

func GetFileURL(bucketName, objectName string) (*url.URL, error) {
	reqParams := make(url.Values)

	preSignedURL, err := setup.Inst.MinioClient.PresignedGetObject(*setup.Inst.MinioContext, bucketName, objectName, time.Hour*24, reqParams)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("failed to get file preSigned url")
		return nil, err
	}

	return preSignedURL, nil
}
