# 🔹 1. Используем официальный образ Go для сборки (с Alpine для меньшего размера)
FROM golang:1.23.4-alpine AS builder

# 🔹 2. Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# 🔹 3. Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# 🔹 4. Копируем исходный код проекта
COPY . .

# 🔹 5. Компилируем бинарник
RUN go build -o /main ./cmd/server/main.go

# 🔹 6. Финальный минимальный образ
FROM alpine:3

# 🔹 7. Устанавливаем зависимости (например, сертификаты для HTTPS)
RUN apk --no-cache add ca-certificates

# 🔹 8. Копируем бинарный файл из builder-образа
COPY --from=builder /main /bin/main

# 🔹 9. Делаем бинарник исполняемым
RUN chmod +x /bin/main

# 🔹 10. Указываем команду для запуска
ENTRYPOINT ["/bin/main"]
