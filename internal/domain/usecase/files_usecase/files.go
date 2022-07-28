package files_usecase

import (
	"aero-internship/gen/api"
	"aero-internship/internal/adapters"
	"aero-internship/internal/domain/entity/files"
	"aero-internship/pkg/config"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
)

type FileService struct {
	cfg          *config.Config
	dataTransfer adapters.DataTransfer
}

func NewFileService(cfg *config.Config, dt adapters.DataTransfer) *FileService {
	return &FileService{cfg: cfg, dataTransfer: dt}
}

func (s FileService) GetFile(ctx context.Context, bucketName, fileName string) (*api.File, error) {

	obj, err := s.dataTransfer.Storage.GetFile(ctx, bucketName, fileName)
	defer obj.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to get file. err: %w", err)
	}

	objectInfo, err := obj.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file. err: %w", err)
	}

	buffer := make([]byte, objectInfo.Size)
	_, err = obj.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to get objects. err: %w", err)
	}

	dateCreate, _ := strconv.Atoi(objectInfo.UserMetadata["DateCreate"])
	encodedBytes := base64.StdEncoding.EncodeToString(buffer)

	f := api.File{
		Id:         objectInfo.Key,
		Name:       objectInfo.UserMetadata["Name"],
		Ext:        objectInfo.UserMetadata["Ext"],
		Base64:     encodedBytes,
		DateCreate: int64(dateCreate),
		UserId:     objectInfo.UserMetadata["UserId"],
	}

	return &f, nil
}

func (s FileService) GetFilesByNewsID(ctx context.Context, newsID string) ([]*api.File, error) {
	objects, err := s.dataTransfer.Storage.GetBucketFiles(ctx, newsID)
	if err != nil {
		return nil, fmt.Errorf("failed to get objects. err: %w", err)
	}
	if len(objects) == 0 {
		return nil, fmt.Errorf("files not found error")
	}

	var files []*api.File
	for _, obj := range objects {
		stat, err := obj.Stat()
		if err != nil {
			logrus.Errorf("failed to get objects. err: %v", err)
			continue
		}
		buffer := make([]byte, stat.Size)
		_, err = obj.Read(buffer)
		if err != nil && err != io.EOF {
			logrus.Errorf("failed to get objects. err: %v", err)
			continue
		}

		encodedBytes := base64.StdEncoding.EncodeToString(buffer)

		dateCreate, _ := strconv.Atoi(stat.UserMetadata["DateCreate"])
		f := api.File{
			Id:         stat.Key,
			Name:       stat.UserMetadata["Name"],
			Ext:        stat.UserMetadata["Ext"],
			Base64:     encodedBytes,
			DateCreate: int64(dateCreate),
			UserId:     stat.UserMetadata["UserId"],
		}
		files = append(files, &f)
		obj.Close()
	}

	return files, nil
}

func (s FileService) CreateFile(ctx context.Context, newsID string, file files.FileDTO) error {
	err := s.dataTransfer.Storage.UploadFile(ctx, file, newsID, bytes.NewBuffer(file.Bytes))
	if err != nil {
		return err
	}
	return nil
}

func (s FileService) DeleteFile(ctx context.Context, newsID, fileId string) error {
	err := s.dataTransfer.Storage.DeleteFile(ctx, newsID, fileId)
	if err != nil {
		return err
	}
	return nil
}
