package storage

import (
	"context"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	BasePath string // например: "./uploads/"
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{BasePath: basePath}
}

func (s *LocalStorage) UploadFile(ctx context.Context, fileBytes []byte, fileName string) (string, error) {
	fullPath := filepath.Join(s.BasePath, fileName)
	if err := os.WriteFile(fullPath, fileBytes, 0644); err != nil {
		return "", err
	}
	return "/uploads/" + fileName, nil // возвращаем путь для file_url
}

func (s *LocalStorage) DeleteFile(ctx context.Context, fileURL string) error {
	fullPath := "." + fileURL // fileURL начинается с /uploads/
	return os.Remove(fullPath)
}
