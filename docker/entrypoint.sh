#!/bin/sh

# Загружаем переменные окружения
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Заменяем переменные окружения в DATABASE_DSN
export DATABASE_DSN="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${POSTGRES_DB}?sslmode=disable"

# Ожидаем, пока PostgreSQL будет доступен
echo "Ожидание запуска PostgreSQL..."
until nc -z -v -w30 $DATABASE_HOST $DATABASE_PORT; do
  echo "Ждем PostgreSQL..."
  sleep 2
done

echo "PostgreSQL доступен, запускаем миграции..."

# Запуск миграций
migrate -path /app/database/migrations -database "$DATABASE_DSN" up

echo "Миграции применены, запускаем сервис..."
exec "$@"
