
---

# Проект: go_wserve

### Микросервис. Работа с чистым кодом

---

## Запуск проекта

Для запуска микросервиса:

```bash
go run cmd/main.go
```

---

## Внедрение зависимостей (Dependency Injection)

Структуры `Service` и `Repository` используют три интерфейса:

- `Authorisation`
- `TodoList`
- `TodoItem`

---

### Архитектура зависимостей:

1. **Handler** зависит от **Service** (уровень выше).
2. **Service** зависит от **Repository** (уровень выше).
3. **Repository** > **Service** > **Handler**

**Service** взаимодействует с базой данных, а **Handler** вызывает методы **Service**.

---

### Подключение postgres:

К `repository` добавлено подключение базы данных postgres. Через файлы **config.yml** и **.env** устанавливается конфигурация для подключения к ней.

Команды выполнения через Docker
```shell
docker run --name=go_verse-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
brew install golang-migrate
976  migrate create -ext sql -dir ./shema -seq init
migrate -path ./shema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
docker exec -it 179bf18a37d7 /bin/bash
psql -U postgres
migrate -path ./shema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down
```
для обновления файла миграции
```shell
psql -U postgres
update schema_migrations set version='000001', dirty=false;
```
для работы с базой данных `psql`
```shell
docker ps
docker exec -it 179bf18a37d7 /bin/bash
psql -U postgres
# select * from users;
```
---

### Регистрация:

Пользователь регистрируется с хешированным паролем, если отправить **POST** запрос с телом:
```json
{
    "name": "namalkin",
    "username": "nanomalkin",
    "password": "qwerty"
}
```