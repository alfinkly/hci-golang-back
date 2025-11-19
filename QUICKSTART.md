# Быстрый старт (Quick Start)

Это руководство поможет вам быстро запустить Pharmacy Backend API.

## Вариант 1: Запуск с Docker Compose (Рекомендуется)

Самый простой способ запустить приложение вместе с PostgreSQL:

```bash
# 1. Клонируйте репозиторий
git clone https://github.com/alfinkly/hci-golang-back.git
cd hci-golang-back

# 2. Создайте .env файл
cp .env.example .env
# Отредактируйте .env если нужно изменить настройки

# 3. Запустите приложение и базу данных
docker-compose up -d

# 4. Проверьте статус
curl http://localhost:8080/health
```

API будет доступен по адресу: `http://localhost:8080`

## Вариант 2: Локальный запуск

Если у вас уже установлен PostgreSQL:

```bash
# 1. Клонируйте репозиторий
git clone https://github.com/alfinkly/hci-golang-back.git
cd hci-golang-back

# 2. Создайте базу данных
createdb pharmacy_db

# 3. Создайте .env файл
cp .env.example .env
# Настройте параметры подключения к PostgreSQL в .env

# 4. Установите зависимости
go mod download

# 5. Запустите приложение
go run main.go
```

## Вариант 3: Только PostgreSQL в Docker

```bash
# 1. Запустите PostgreSQL в Docker
docker-compose up -d postgres

# 2. Создайте .env файл
cp .env.example .env

# 3. Запустите приложение
go run main.go
```

## Первые шаги

### 1. Проверьте работу API

```bash
curl http://localhost:8080/health
```

Ожидаемый ответ:
```json
{
  "status": "ok",
  "message": "Pharmacy API is running"
}
```

### 2. Зарегистрируйте пользователя

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "email": "admin@pharmacy.com",
    "password": "admin123",
    "role": "admin"
  }'
```

### 3. Войдите и получите токен

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

Сохраните полученный `token` для дальнейших запросов.

### 4. Используйте токен для защищенных маршрутов

```bash
# Замените YOUR_TOKEN на полученный токен
curl http://localhost:8080/api/medicines \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Автоматическое тестирование

Используйте скрипт для автоматического тестирования всех endpoints:

```bash
./test_api.sh
```

Скрипт требует установленный `jq` для форматирования JSON:
```bash
# Ubuntu/Debian
sudo apt-get install jq

# macOS
brew install jq
```

## Полезные команды

```bash
# Сборка приложения
make build

# Запуск приложения
make run

# Очистка скомпилированных файлов
make clean

# Обновление зависимостей
make deps

# Форматирование кода
go fmt ./...

# Проверка кода
go vet ./...

# Запуск тестов
go test ./... -v
```

## Остановка приложения

### Docker Compose
```bash
docker-compose down
```

### Локальный запуск
Нажмите `Ctrl+C` в терминале где запущено приложение

## Следующие шаги

- Прочитайте [README.md](README.md) для полной документации API
- Изучите примеры запросов для всех endpoints
- Настройте параметры в `.env` под свои нужды
- Интегрируйте API с вашим фронтендом

## Возможные проблемы

### Не удается подключиться к базе данных

1. Проверьте, что PostgreSQL запущен
2. Проверьте настройки в `.env`
3. Убедитесь, что база данных `pharmacy_db` создана

### Порт 8080 уже занят

Измените `PORT` в `.env` файле на другой порт, например `3000`.

### Ошибки при компиляции

Убедитесь, что используется Go 1.25 или выше:
```bash
go version
```

## Поддержка

Если у вас возникли вопросы или проблемы, создайте issue в GitHub репозитории.
