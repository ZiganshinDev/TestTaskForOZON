# TestTaskForOZON by ZiganshinDev

## Как запустить:

1. Склонировать репозиторий
2. Запустить Docker Descktop 
3. Открыть директорию с проектом
4. Собрать проект (docker-compose up --build) 

### (Опционально) Чтобы сменить хранилище: 

Измените .env файл STORAGE_TYPE=in-memory/postgres (на своё усмотрение) 

postgres - по умолчанию

## Описание задачи: 

Необходимо реализовать сервис, который должен предоставлять API по созданию сокращенных ссылок следующего формата:

Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться только одна сокращенная ссылка.
Ссылка должна быть длинной 10 символов
Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание)
Сервис должен быть написан на Go и принимать следующие запросы по http:

Метод Post, который будет сохранять оригинальный URL в базе и возвращать сокращённый
Метод Get, который будет принимать сокращённый URL и возвращать оригинальный URL
Решение должно быть предоставлено в «конечном виде», а именно:

Сервис должен быть распространён в виде Docker-образа
В качестве хранилища ожидается использовать in-memory решение И postgresql. Какое хранилище использовать указывается параметром при запуске сервиса.
Покрыть реализованный функционал Unit-тестами

## Демонстрация работы:

![Imgur](https://i.imgur.com/VPz9zKy.png)

![Imgur](https://i.imgur.com/nMyP5Kb.png)

## API:

1. Тело запроса/ответа - в формате JSON.
2. В случае ошибки возвращается необходимый HTTP код, в теле содержится описание ошибки

## Реализация:

1. Следование дизайну REST API.
2. Подход "Чистой Архитектуры" и техника внедрения зависимости.
3. Работа с фреймворком gorilla/mux.
4. Работа с БД Postgres с использованием библиотеки lib/pq и написанием SQL запросов.
5. Запуск из Docker.