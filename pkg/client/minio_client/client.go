package minio_client

import (
	"aero-internship/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

func NewMinioClient(cfg *config.Config) (*minio.Client, error) {

	minioClient, err := minio.New(cfg.GetMinioEndpoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.GetMinioAccessKeyId(), cfg.GetMinioSecretAccessKey(), ""),
		Secure: false,
	})
	if err != nil {
		logrus.Errorf("Error occured during new minio client creation: %s", err)
		return nil, err
	}

	return minioClient, nil
}
