package storage

import (
	"github.com/minio/minio-go/v7"
	"log"
	"os"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Client *minio.Client

func InitMinIO() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"

	var err error
	Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("❌ MinIO init failed: %v", err)
	}
	log.Println("✅ MinIO connected:", endpoint)
}
