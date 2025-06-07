package config

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/raxaris/ipromise-backend/internal/storage"
	"log"
	"os"
)

func NewMinIOClient() (*minio.Client, string, string, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"
	bucket := os.Getenv("MINIO_BUCKET")
	publicEndpoint := os.Getenv("MINIO_PUBLIC_URL")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, "", "", err
	}

	log.Printf("✅ MinIO connected to SDK: %s | public URL: %s", endpoint, publicEndpoint)
	return client, bucket, publicEndpoint, nil
}

func CreateStorage() storage.Storage {
	minioClient, bucketName, minioPublicURL, err := NewMinIOClient()
	if err != nil {
		return nil
	}

	err = EnsureMinioBucket(minioClient, bucketName)
	if err != nil {
		log.Println("❌ Failed to ensure bucket:", err)
		return nil
	}

	store := storage.NewMinIOStorage(minioClient, bucketName, minioPublicURL)
	return store
}

func EnsureMinioBucket(client *minio.Client, bucketName string) error {
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
		log.Println("✅ Bucket created:", bucketName)
	} else {
		log.Println("✅ Bucket already exists:", bucketName)
	}
	return nil
}
