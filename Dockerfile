FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /ipromise-backend cmd/server/main.go

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /ipromise-backend /bin/ipromise-backend

COPY .env .env

CMD ["/bin/ipromise-backend"]
