# Pharmacy Backend API

Backend API для аптечной базы данных, построенный на Go с использованием Fiber v3, sqlx и JWT авторизации.

## Технологии

- **Go 1.25+**
- **Fiber v3** - веб-фреймворк
- **sqlx** - расширение для работы с SQL
- **PostgreSQL** - база данных
- **JWT** - авторизация
- **bcrypt** - хеширование паролей

## Возможности

- ✅ JWT авторизация
- ✅ CRUD операции для лекарств
- ✅ CRUD операции для поставщиков
- ✅ Управление закупками
- ✅ Управление продажами
- ✅ Middleware для всех защищенных маршрутов
- ✅ Автоматическое обновление количества при закупках/продажах
- ✅ CORS поддержка

## Структура проекта

```
.
├── config/          # Конфигурация приложения
├── database/        # Подключение к БД и инициализация схемы
├── handlers/        # HTTP обработчики
├── middleware/      # Middleware функции
├── models/          # Модели данных
├── utils/           # Утилиты (JWT, bcrypt)
├── main.go          # Точка входа
├── .env.example     # Пример конфигурации
└── go.mod           # Go зависимости
```

## Установка и запуск

### Предварительные требования

- Go 1.25 или выше
- PostgreSQL 12 или выше

### Шаги установки

1. Клонируйте репозиторий:
```bash
git clone https://github.com/alfinkly/hci-golang-back.git
cd hci-golang-back
```

2. Установите зависимости:
```bash
go mod download
```

3. Создайте базу данных PostgreSQL:
```sql
CREATE DATABASE pharmacy_db;
```

4. Создайте файл `.env` на основе `.env.example`:
```bash
cp .env.example .env
```

5. Настройте переменные окружения в `.env`:
```env
PORT=8080
APP_ENV=development

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=pharmacy_db
DB_SSLMODE=disable

JWT_SECRET=your-secret-key-change-this
JWT_EXPIRATION=24h
```

6. Запустите приложение:
```bash
go run main.go
```

Сервер запустится на `http://localhost:8080`

## API Endpoints

### Публичные маршруты

#### Регистрация
```http
POST /api/auth/register
Content-Type: application/json

{
  "username": "user",
  "email": "user@example.com",
  "password": "password123",
  "role": "user"
}
```

#### Вход
```http
POST /api/auth/login
Content-Type: application/json

{
  "username": "user",
  "password": "password123"
}
```

### Защищенные маршруты (требуется JWT токен)

Все защищенные маршруты требуют заголовок:
```
Authorization: Bearer <your-jwt-token>
```

#### Профиль пользователя
```http
GET /api/profile
```

#### Лекарства

```http
GET    /api/medicines        # Получить все лекарства
GET    /api/medicines/:id    # Получить лекарство по ID
POST   /api/medicines        # Создать лекарство
PUT    /api/medicines/:id    # Обновить лекарство
DELETE /api/medicines/:id    # Удалить лекарство
```

Пример создания лекарства:
```json
{
  "name": "Аспирин",
  "description": "Обезболивающее средство",
  "manufacturer": "Bayer",
  "price": 150.50,
  "quantity": 100,
  "expiry_date": "2025-12-31T00:00:00Z",
  "category": "Обезболивающие",
  "requires_prescription": false
}
```

#### Поставщики

```http
GET    /api/suppliers        # Получить всех поставщиков
GET    /api/suppliers/:id    # Получить поставщика по ID
POST   /api/suppliers        # Создать поставщика
PUT    /api/suppliers/:id    # Обновить поставщика
DELETE /api/suppliers/:id    # Удалить поставщика
```

Пример создания поставщика:
```json
{
  "name": "ООО Фарма",
  "contact_person": "Иван Иванов",
  "phone": "+7 (999) 123-45-67",
  "email": "contact@pharma.com",
  "address": "Москва, ул. Примерная, д. 1"
}
```

#### Закупки

```http
GET    /api/purchases        # Получить все закупки
GET    /api/purchases/:id    # Получить закупку по ID
POST   /api/purchases        # Создать закупку
DELETE /api/purchases/:id    # Удалить закупку
```

Пример создания закупки:
```json
{
  "medicine_id": 1,
  "supplier_id": 1,
  "quantity": 50,
  "unit_price": 120.00
}
```

#### Продажи

```http
GET    /api/sales            # Получить все продажи
GET    /api/sales/:id        # Получить продажу по ID
POST   /api/sales            # Создать продажу
DELETE /api/sales/:id        # Удалить продажу
```

Пример создания продажи:
```json
{
  "medicine_id": 1,
  "quantity": 2
}
```

#### Health Check
```http
GET /health
```

## База данных

Схема базы данных автоматически создается при первом запуске приложения.

### Таблицы:

- **users** - пользователи системы
- **medicines** - лекарства
- **suppliers** - поставщики
- **purchases** - закупки
- **sales** - продажи

## Middleware

Все защищенные маршруты используют следующие middleware:

1. **CORSMiddleware** - обработка CORS запросов
2. **LoggingMiddleware** - логирование запросов
3. **JWTMiddleware** - проверка JWT токенов

## Лицензия

MIT License - см. файл LICENSE для деталей