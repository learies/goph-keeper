# go-keeper

Разработана клиент-серверная система, позволяющая пользователю надёжно и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.

# Проектирование решения

Проработка решения:

![Основная схема](info/main-scheme.png)

Клиент. Ретраи:

![Проектирование ретраев](info/client-retry.png)

Шифрование на всех уровнях:

![Шифрование](info/main-encryption.png)

Синхронизация с одним сервером:

![Синхронизация](info/sync-one-server.png)

Синхронизация с несколькими серверами:

![Синхронизация 2](info/sync-many-servers.png)

Чистая архитектура на клиенте и сервере:

![Чистая архитектура](info/main-app-architecture.png)

# Покрытие тестами

![Юнит-тесты](info/tests-coverage.png)

# Что применял?

- [x] Регистрация и аутентификация пользователей
- [x] Толстый консольный клиент
- [x] gRPC, включая логирование запросов и ответов
- [x] Шифрование на клиенте и сервере
- [x] Сервисный слой
- [x] Покрытие юнит-тестами 80% состояний
- [x] Описание экспортированных функций, типов, переменных, а также пакетов системы
