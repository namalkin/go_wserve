
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
