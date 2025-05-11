package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/raxaris/ipromise-backend/internal/storage"
)

func main() {
	// Создаём локальный сторедж
	localStorage := storage.NewLocalStorage("./uploads/")

	// Файл, который будем сохранять (например, тестовая картинка)
	fileBytes := []byte("Hello, this is a test file!")
	fileName := "testfile.txt"

	// Пробуем сохранить файл
	url, err := localStorage.UploadFile(context.Background(), fileBytes, fileName)
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
	}

	fmt.Printf("File saved successfully at: %s\n", url)

	// Проверяем, что файл существует
	if _, err := os.Stat("." + url); os.IsNotExist(err) {
		log.Fatalf("file does not exist at path: %s", url)
	} else {
		fmt.Println("File exists, upload worked correctly!")
	}

	// Теперь пробуем удалить файл
	if err := localStorage.DeleteFile(context.Background(), url); err != nil {
		log.Fatalf("failed to delete file: %v", err)
	} else {
		fmt.Println("File deleted successfully!")
	}
}
