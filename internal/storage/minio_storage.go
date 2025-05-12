package storage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"mime"
	"path/filepath"
)

type minioStorage struct {
	client     *minio.Client
	bucketName string
	endpoint   string // для формирования публичного URL
}

func NewMinIOStorage(client *minio.Client, bucket string, endpoint string) Storage {
	return &minioStorage{
		client:     client,
		bucketName: bucket,
		endpoint:   endpoint,
	}
}

func (s *minioStorage) UploadFile(ctx context.Context, fileBytes []byte, fileName string) (string, error) {
	ext := filepath.Ext(fileName)
	contentType := mime.TypeByExtension(ext)

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	objectName := uuid.New().String() + ext

	_, err := s.client.PutObject(ctx, s.bucketName, objectName, bytes.NewReader(fileBytes), int64(len(fileBytes)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("http://%s/%s/%s", s.endpoint, s.bucketName, objectName)
	return url, nil
}

func (s *minioStorage) DeleteFile(ctx context.Context, fileURL string) error {
	_, objectName := filepath.Split(fileURL)
	return s.client.RemoveObject(ctx, s.bucketName, objectName, minio.RemoveObjectOptions{})
}
