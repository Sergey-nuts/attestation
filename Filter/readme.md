# Сервис цензуры комментариев


## Методы серваиса Filter:

Method POST: `<host>:2020/filter`

request JSON:

```json
{
    "text": "HA - HA"
}
```

## Сборка докер образа

сбор doccker образа
`make build`

Запуск докер контейнера
`make run`