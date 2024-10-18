# Simple CRUD

## Task 1
Разработать простой круд на 2 эндпоинта - создание таски и получение его по id.

1) Для реализации поднять постгрес на docker. 
2) Использовать дефолтные библиотеки go для работы с сетью и бд.
3) Использовать git только через cmd
4) Постараться выполнить за 1 час

**RESULT**: 2 часа, пришлось решать проблему с подключением к бд - устанавливать локально pg, изучать некоторые вещи. 

## Task 2
Прикрутить миграции, чтобы при каждом запуске приложения - накатывалась структура бд и заполнялось бд некоторыми значениями, а после остановки контейнера - удалялось.

1) Миграции накатываются
2) Сделать за час

**RESULT**: 1 час, использовал Makefile и migration-go, но миграции с инсертами значений в последнем почему-то не работают, так что без инсертов - чисто структура.

**RESULT NEW**: ещё 1 час потратил на настройку .env и подключения переменных окружения оттуда в docker-compose и makefile для приватной инфы, а также добавил в makefile новые штуки - типа sleep и поднятие доккера на запуске, а после накат миграций и тд.

## GETTING START
Start app - up docker:
    ```make run```
Clean and stop docker:
    ```make clean```

