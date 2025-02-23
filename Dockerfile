# Базовый образ с Go
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарный файл
RUN go build -o /ipromise-backend cmd/server/main.go

# Финальный образ
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /ipromise-backend /bin/ipromise-backend

# Копируем .env (если потребуется)
COPY .env .env

# Запуск приложения
CMD ["/bin/ipromise-backend"]
