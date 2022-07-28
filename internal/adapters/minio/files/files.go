package files

import (
	"aero-internship/internal/domain/entity/files"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
	"time"
)

type FileStorage struct {
	minioClient *minio.Client
}

func NewFilesStorage(minioClient *minio.Client) *FileStorage {
	return &FileStorage{minioClient: minioClient}
}

func (c *FileStorage) GetFile(ctx context.Context, bucketName, fileId string) (*minio.Object, error) {
	reqCtx, _ := context.WithTimeout(ctx, 10*time.Second)

	obj, err := c.minioClient.GetObject(reqCtx, bucketName, fileId, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get file with id: %s from minio bucket %s. err: %w", fileId, bucketName, err)
	}

	return obj, nil
}

func (c *FileStorage) GetBucketFiles(ctx context.Context, bucketName string) ([]*minio.Object, error) {
	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var files []*minio.Object
	for lobj := range c.minioClient.ListObjects(reqCtx, bucketName, minio.ListObjectsOptions{WithMetadata: true}) {
		if lobj.Err != nil {
			logrus.Errorf("failed to list object from minio bucket %s. err: %v", bucketName, lobj.Err)
			continue
		}
		object, err := c.minioClient.GetObject(ctx, bucketName, lobj.Key, minio.GetObjectOptions{})
		if err != nil {
			logrus.Errorf("failed to get object key=%s from minio bucket %s. err: %v", lobj.Key, bucketName, lobj.Err)
			continue
		}
		files = append(files, object)
	}
	return files, nil
}

func (c *FileStorage) UploadFile(ctx context.Context, fileDTO files.FileDTO, bucketName string, reader io.Reader) error {
	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	exists, errBucketExists := c.minioClient.BucketExists(ctx, bucketName)
	if errBucketExists != nil || !exists {
		logrus.Printf("no bucket %s. creating new one...", bucketName)
		err := c.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create new bucket. err: %w", err)
		}
	}

	logrus.Printf("put new object %s to bucket %s", fileDTO.Name, bucketName)
	_, err := c.minioClient.PutObject(reqCtx, bucketName, fileDTO.Name, reader, fileDTO.FileSize,
		minio.PutObjectOptions{
			UserMetadata: map[string]string{
				"Name":       fileDTO.Name,
				"DateCreate": strconv.FormatInt(fileDTO.DateCreate, 10),
				"Ext":        fileDTO.Ext,
				"UserID":     fileDTO.UserId,
			},
			ContentType: "application/octet-stream",
		})
	if err != nil {
		return fmt.Errorf("failed to upload file. err: %w", err)
	}
	return nil
}

func (c *FileStorage) DeleteFile(ctx context.Context, bucketName, fileName string) error {
	err := c.minioClient.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file. err: %w", err)
	}
	return nil
}
