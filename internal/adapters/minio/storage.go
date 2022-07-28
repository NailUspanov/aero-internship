package minio

import (
	"aero-internship/internal/adapters/minio/files"
	entity_files "aero-internship/internal/domain/entity/files"
	"github.com/minio/minio-go/v7"

	"context"
	"io"
)

type FilesStorage interface {
	GetFile(ctx context.Context, bucketName, fileId string) (*minio.Object, error)
	GetBucketFiles(ctx context.Context, bucketName string) ([]*minio.Object, error)
	UploadFile(ctx context.Context, fileDTO entity_files.FileDTO, bucketName string, reader io.Reader) error
	DeleteFile(ctx context.Context, bucketName, fileName string) error
}

type Storage struct {
	FilesStorage
}

func NewStorage(client *minio.Client) *Storage {
	return &Storage{FilesStorage: files.NewFilesStorage(client)}
}
