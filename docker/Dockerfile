# Используем Golang + Alpine
FROM golang:alpine AS builder

# Устанавливаем нужные пакеты
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

# Устанавливаем migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    
# Создаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY app .

# Загружаем зависимости
RUN go mod tidy

# Собираем Go-приложение (если main.go в cmd/, поменяй путь)
RUN CGO_ENABLED=0 GOOS=linux go build -o /cmd/main ./cmd

# Копируем entrypoint
COPY docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Открываем порт
EXPOSE 8080

# Запускаем entrypoint
ENTRYPOINT ["/entrypoint.sh"]

# Запускаем приложение
CMD ["/cmd/main"]
