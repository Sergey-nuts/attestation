# Сервис комментариев

## env

необязательный env для смены хоста сервиса фильтрации

filter=yourFilterHost

## Методы сервиса Comments:

Method GET: `<host>:2010/comments/<postid>` 
Получить комментарии к новости с идентификатором postid


Method POST: `<host>:2010/comments`
Добавить коммнетарий к новости с идентификатором postid

request JSON:
```json
{
    "postid": 10,
    "text": "HA - HA",
    "author": "some_user",
    "pubtime": "2023-09-22T13:07:08+03:00"
}
```

## Сборка докер образа

сбор doccker образа
`make build`

Запуск докер контейнера
`make run`