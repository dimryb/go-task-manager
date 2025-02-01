# go-task-manager

# Документация по сервису управления задачами с аналитикой

## Оглавление
1. [Введение](#введение)
2. [Запуск приложения через Docker-compose](#запуск-приложения-через-docker-compose)
3. [API-запросы](#api-запросы)
4. [Примеры JSON для импорта/экспорта задач](#примеры-json-для-импортаэкспорта-задач)
5. [Swagger-документация](#swagger-документация)
6. [Тестирование](#тестирование)
7. [Логирование](#логирование)

## Введение

Данный сервис представляет собой веб-приложение для управления задачами с расширенными функциями аналитики. Сервис реализован на языке Golang с использованием PostgreSQL в качестве базы данных. Архитектура приложения построена по принципу Clean Architecture.

## Запуск приложения через Docker-compose

Для запуска приложения необходимо выполнить следующие шаги:

1. Убедитесь, что у вас установлены Docker и Docker-compose.
2. Скачайте исходный код проекта.
3. Перейдите в директорию проекта.
4. Создайте файл `.env` в корневой директории проекта и заполните его необходимыми переменными окружения на основе `.env.example`:

    ```env
    DATABASE_URL=postgres://postgres:123456@postgres:5432/go_task_manager
    LOG_LEVEL=debug
    HTTP_PORT=8080
    ```

5. Запустите приложение с помощью команды:

    ```bash
    docker-compose up --build
    ```

6. После успешного запуска приложение будет доступно по адресу `http://localhost:8080`.

## API-запросы

### Регистрация пользователя

**Запрос:**

```bash
curl -X 'POST' \
  'http://localhost:8080/auth/register' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "user",
    "password": "password"
}'
```

**Ответ:**

```json
{
    "ok": true
}
```

### Вход и получение JWT

**Запрос:**

```bash
curl -X 'POST' \
  'http://localhost:8080/auth/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "user",
    "password": "password"
}'
```

**Ответ:**

```json
{
    "ok": true,
    "result": "your_jwt_token"
}
```

### Получение списка задач

**Запрос:**

```bash
curl -X 'GET' \
  'http://localhost:8080/tasks?status=pending&priority=medium&due_date=2020-01-02T12%3A00%3A00Z&title=Title' \
  -H 'accept: application/json' \
  -H "Authorization: Bearer your_jwt_token"
```

**Ответ:**

```json
{
   "ok": true,
   "result": [
      {
         "id": 1,
         "title": "Title",
         "description": "Description",
         "status": "pending",
         "priority": "medium",
         "due_date": "2020-01-02T12:00:00Z",
         "created_at": "2020-01-02T09:30:57Z",
         "updated_at": "2020-01-02T09:30:57Z"
      },
      {
         "id": 2,
         "title": "Title2",
         "description": "Description",
         "status": "pending",
         "priority": "medium",
         "due_date": "2020-01-02T12:00:00Z",
         "created_at": "2020-01-02T09:37:30Z",
         "updated_at": "2020-01-02T09:37:30Z"
      }
   ]
}
```

### Добавление задачи

**Запрос:**

```bash
curl -X 'POST' \
  'http://localhost:8080/tasks' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer your_jwt_token' \
  -H 'Content-Type: application/json' \
  -d '{
  "description": "Description",
  "due_date": "2020-01-02T12:00:00Z",
  "priority": "medium",
  "status": "pending",
  "title": "Title"
}'
```

**Ответ:**

```json
{
   "ok": true,
   "result": 2
}
```

### Обновление задачи

**Запрос:**

```bash
curl -X 'PUT' \
  'http://localhost:8080/tasks/1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer your_jwt_token' \
  -H 'Content-Type: application/json' \
  -d '{
  "description": "Updated Description",
  "due_date": "2025-02-01T15:00:00Z",
  "priority": "high",
  "status": "in_progress",
  "title": "Updated Title"
}'
```

**Ответ:**

```json
{
    "ok": true
}
```

### Удаление задачи

**Запрос:**

```bash
curl -X 'DELETE' \
  'http://localhost:8080/tasks/1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer your_jwt_token'
```

**Ответ:**

```json
{
   "ok": true,
   "result": 1
}
```

### Экспорт задач в JSON

**Запрос:**

```bash
curl -X 'GET' \
  'http://localhost:8080/tasks/export' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer your_jwt_token' \
  --output tasks.json
```

**Ответ:**

Файл tasks.json будет скачан на ваш компьютер. Пример содержимого файла:

```json
[
    {
        "id": 1,
        "title": "Task 1",
        "description": "Description 1",
        "status": "pending",
        "priority": "medium",
        "due_date": "2023-10-01T00:00:00Z",
        "created_at": "2023-09-25T12:00:00Z",
        "updated_at": "2023-09-25T12:00:00Z"
    }
]
```

### Импорт задач из JSON

**Запрос:**

```bash
curl -X 'POST' \
  'http://localhost:8080/tasks/import' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer your_jwt_token' \
  -H 'Content-Type: multipart/form-data' \
  -F 'file=@tasks.json;type=application/json'
```

**Ответ:**

```json
{
   "ok": true,
   "result": "Tasks imported successfully"
}
```

## Примеры JSON для импорта/экспорта задач

### Пример файла JSON для импорта задач

```json
[
    {
        "id": 1,
        "title": "Imported Task 1",
        "description": "Imported Description 1",
        "status": "pending",
        "priority": "low",
        "due_date": "2023-10-15T00:00:00Z",
        "created_at": "2023-09-25T12:00:00Z",
        "updated_at": "2023-09-25T12:00:00Z"
    },
    {
        "id": 2,
        "title": "Imported Task 2",
        "description": "Imported Description 2",
        "status": "in_progress",
        "priority": "medium",
        "due_date": "2023-10-20T00:00:00Z",
        "created_at": "2023-09-25T12:00:00Z",
        "updated_at": "2023-09-25T12:00:00Z"
    }
]
```

### Пример файла JSON для экспорта задач

```json
[
    {
        "id": 1,
        "title": "Task 1",
        "description": "Description 1",
        "status": "pending",
        "priority": "medium",
        "due_date": "2023-10-01T00:00:00Z",
        "created_at": "2023-09-25T12:00:00Z",
        "updated_at": "2023-09-25T12:00:00Z"
    },
    {
        "id": 2,
        "title": "New Task",
        "description": "New Description",
        "status": "in_progress",
        "priority": "high",
        "due_date": "2023-10-10T00:00:00Z",
        "created_at": "2023-09-25T12:05:00Z",
        "updated_at": "2023-09-25T12:10:00Z"
    }
]
```

## Swagger-документация

Swagger-документация доступна по адресу `http://localhost:8080/swagger/index.html` после запуска приложения. Она предоставляет полное описание всех доступных API-методов, их параметров и примеров запросов/ответов.

## Тестирование

### Unit-тесты

Для запуска unit-тестов выполните следующую команду:

```bash
go test ./...
```

### Запуск интеграционных тестов в Docker
1. Создайте файл .env.test в корневой директории проекта и заполните его необходимыми переменными окружения на основе .env.example:
```dotenv
DATABASE_URL=postgres://postgres:123456@postgres_test:5432/go_task_manager_test
LOG_LEVEL=debug
HTTP_PORT=8080
```
2. Запустите тесты с помощью команды:
```bash
docker-compose -f docker-compose.test.yaml up --build --abort-on-container-exit
```
3. После завершения тестов контейнеры будут остановлены, и вы увидите результаты выполнения тестов.

## Логирование

Логирование реализовано с использованием библиотеки `logrus`. Логи выводятся в консоль и сохраняются в файл `app.log`.
