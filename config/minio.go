package config

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/raxaris/ipromise-backend/internal/storage"
	"log"
	"os"
)

// NewMinIOClient инициализирует и возвращает *minio.Client
func NewMinIOClient() (*minio.Client, string, string, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT") // например: localhost:9000
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"
	bucket := os.Getenv("MINIO_BUCKET")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, "", "", err
	}

	log.Println("✅ MinIO connected:", endpoint)
	return client, bucket, endpoint, nil
}

func CreateStorage() storage.Storage {
	minioClient, bucketName, minioEndpoint, err := NewMinIOClient()
	if err != nil {
		return nil
	}

	store := storage.NewMinIOStorage(minioClient, bucketName, minioEndpoint)
	return store
}
