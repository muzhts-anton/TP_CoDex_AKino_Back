# 2022_1_CoDex
Кинопоиск

## Деплой

Front: https://tp-frontkinopoisk.herokuapp.com/
Back: https://teamprojectkinopoisk.herokuapp.com/

## API

https://app.swaggerhub.com/apis/CoDex9/CoDex/1.0.0#/info

## Авторы

* [**Антон Мужецкий**](https://github.com/muzhts-anton) - *Бэкенд*
* [**Киселёв Виктор**](https://github.com/Kislv)        -  *Бэкенд*
* [**Калашков Павел**](https://github.com/kalashkovpaul) - *Фронтенд*
* [**Костинич Константин**](https://github.com/Kostich31) - *Фронтенд*

## Менторы
* [**Клонов Александр**](https://github.com/Shureks-den)      - *Фронтенд*
* [**Богомолова Мария**](https://github.com/keithzetterstrom) - *Бэкенд*
* [**Комиссаров Николай**](https://www.youtube.com/watch?v=dQw4w9WgXcQ) - *Интерфейсы*

## Ссылка на front-end

https://github.com/frontend-park-mail-ru/2022_1_CoDex/tree/develop

Команда для запуска тестов: go test -coverpkg=./... -coverprofile=cover ./... && cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out && go tool cover -func=cover.out


РК:
1. 
2. Сделать еще одну структуру для хранения пользователя