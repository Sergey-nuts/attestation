# Сервис APIGateway

## методы сервиса APIGateway:

Method GET: `<host>:8080/news?page`
Получить список новостей 


Method GET: `<host>:8080/news/<postid>/full`
Получить детальную информацию о новости и комментариях к ней


Method POST: `<host>:8080/news/comments`
Добавить коммнетарий к новости с идентификатором postid

request JSON:
```json
{
    "Text": "Ha-Ha-Ha this is very funny",
    "Author": "user",
    "PubTime": "2023-09-22T15:28:38+03:00"
}
```

Method GET `<host>:8080/news/search?value`
поиск новостей по названию