FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o file-transfer-example

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/file-transfer-example .

CMD ["./file-transfer-example", "server"]
