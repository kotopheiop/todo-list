# Todo List App [![Go version](https://img.shields.io/badge/go-1.21-blue.svg)](https://golang.org/dl/) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/kotopheiop/todo-list/master/LICENSE) 


Простое приложение для списка дел, созданное с помощью Go, Vue и Redis.

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
Соберите и запустите приложение с помощью Docker Compose:
```
docker-compose up -d --build
```

Приложение будет доступно по адресу http://localhost:8080.