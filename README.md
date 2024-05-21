# Todo List App [![Go version](https://img.shields.io/badge/go-1.21-blue.svg)](https://golang.org/dl/) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/kotopheiop/todo-list/master/LICENSE) 


Простое приложение для списка дел, созданное с помощью Go и Vue.

## Особенности

- Создавайте, редактируйте и удаляйте задачи
- Отмечайте задачи как выполненные или нет
- Фильтруйте задачи по статусу

## Установка

Для запуска приложения вам нужно иметь Docker установленный на вашем компьютере. Вы можете скачать Docker [здесь](https://www.docker.com/get-started).

Клонируйте репозиторий на ваш локальный компьютер:
```bash
git clone https://github.com/kotopheiop/todo-list.git
```
Перейдите в директорию проекта:
```
cd todo-list
```
## Запуск

У вас есть два варианта запуска приложения с различными конфигурациями базы данных с использованием Makefile.

### Запуск с MySQL
Для запуска приложения с использованием MySQL, выполните команду:

```bash
make up-mysql
```

### Запуск с MySQL и phpMyAdmin (debug)
Для запуска приложения с использованием MySQL и phpMyAdmin для отладки, выполните команду:

```bash
make up-mysql-debug
```
phpMyAdmin будет доступен по адресу http://localhost:8081.

### Запуск с Redis
Для запуска приложения с использованием Redis, выполните команду:

```bash
make up-redis
```

Приложение будет доступно по адресу http://localhost:8080.