# Используем официальный образ Go для сборки
FROM golang:1.23.0 as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum из корневой директории проекта
COPY ./go.mod ./go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем код PaymentService
COPY . ./

# Переход в директорию PaymentService для сборки
WORKDIR /app/PaymentService/cmd

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o payment-service .

# Используем легкий образ для выполнения
FROM alpine:latest

WORKDIR /root/

# Копируем скомпилированный бинарный файл из builder
COPY --from=builder /app/PaymentService/cmd/payment-service .

# Запускаем приложение
CMD ["./payment-service"]
