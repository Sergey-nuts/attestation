# Задача к итоговой аттестации

В нашем конкретном случае API Gateway будет предоставлять интерфейс для чтения новостей и комментариев к ним, а также позволит оставлять комментарии к новостям.

## Сервис новостей

Данный сервис будет состоять из двух частей:

1. загрузчик обновлений из RSS-канала,
2. API для отдачи новостей по запросу.

В дальнейшем такой сервис, скорее всего, также будет разделён на 2 сервиса, так как количество страниц, которые нужно парсить и количество пользователей — это никак не связанные, независимые величины. С целью уменьшения количества кода мы остановимся на варианте, когда эти функции объединены.

## Сервис комментариев

Будет также состоять из двух частей:

1. механизм автоматической проверки комментариев (модерация): асинхронно запускаем проверку для следующих N комментариев, которые ещё не проверяли,
2. API для запоминания комментариев и метод для отдачи комментариев.