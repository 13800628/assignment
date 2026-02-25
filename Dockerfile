# ビルド環境
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

# 実行環境
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY .env.example .env

EXPOSE 8080
CMD ["./main"]