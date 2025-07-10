# Etapa 1: Construcción con Go
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila el binario con optimización para producción
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w" ./cmd/api

# Etapa 2: Imagen final sin Alpine
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x /app/main

# Expone el puerto en el que corre Fiber
EXPOSE 8080

# Comando para iniciar la app
CMD ["./main"]
