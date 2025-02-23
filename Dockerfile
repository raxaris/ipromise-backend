# 1. Используем официальный образ Go
FROM golang:1.21 AS builder

# 2. Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# 3. Копируем `go.mod` и `go.sum`, скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# 4. Копируем весь код в контейнер
COPY . .

# 5. Собираем бинарник
RUN go build -o main ./cmd/main.go

# 6. Создаём минимальный образ для продакшена
FROM alpine:latest

# 7. Устанавливаем зависимости (например, PostgreSQL-клиент)
RUN apk --no-cache add ca-certificates

# 8. Устанавливаем рабочую директорию и копируем бинарник
WORKDIR /root/
COPY --from=builder /app/main .

# 9. Указываем команду для запуска
CMD ["./main"]
