package storage

import "context"

type Storage interface {
	UploadFile(ctx context.Context, fileBytes []byte, fileName string) (string, error)
	DeleteFile(ctx context.Context, fileURL string) error
}
