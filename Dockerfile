# Etapa de build
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o app ./cmd/api

# Etapa final
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]