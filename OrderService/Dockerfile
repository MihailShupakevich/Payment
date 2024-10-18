# Используем официальный образ Go для сборки
FROM golang:1.23.0 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum из корневой директории проекта
COPY ./go.mod ./go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем код OrderService
COPY ./OrderService ./OrderService

# Переход в директорию OrderService для сборки
WORKDIR /app/OrderService/cmd

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o order-service .

# Используем легкий образ для выполнения
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарный файл из builder
COPY --from=builder /app/OrderService/cmd/order-service .

# Открываем порт 8080
EXPOSE 8080

# Запускаем приложение
CMD ["./order-service"]
