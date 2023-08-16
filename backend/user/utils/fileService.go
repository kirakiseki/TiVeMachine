package utils

import (
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"time"
	"user/setup"
)

func UploadFile(bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	file, err := setup.Inst.MinioClient.PutObject(setup.Inst.MinioContext, bucketName, objectName, reader, objectsize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("failed to upload file")
		return err
	}
	setup.Inst.Logger.Info().Str("bucketName", bucketName).Str("objectName", objectName).Int64("size", file.Size).Msg("Uploaded file")
	return nil
}

func GetFileURL(bucketName, objectName string) (*url.URL, error) {
	reqParams := make(url.Values)

	preSignedURL, err := setup.Inst.MinioClient.PresignedGetObject(setup.Inst.MinioContext, bucketName, objectName, time.Hour*24, reqParams)
	if err != nil {
		setup.Inst.Logger.Error().Err(err).Msg("failed to get file preSigned url")
		return nil, err
	}

	return preSignedURL, nil
}
