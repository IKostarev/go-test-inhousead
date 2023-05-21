# Тестовое задание на позицию Golang Developer at InHouseAd

## Краткое описание приложения:
```
Нужно написать программу, которая будет проверять список сайтов на доступность.
Раз в минуту нужно проверять доступны ли сайты из списка и засекать время доступа к ним.
Есть большое кол-во пользователей, которые хотят знать время доступа к сайтам.
```

## Ендпоинты:
```
1. Получить время доступа к определенному сайту. /{id}, 
где id - полное название сайта, например: /microsoft.com
```
```
2. Получить имя сайта с минимальным временем доступа. /low
```
```
3. Получить имя сайта с максимальным временем доступа. /slow
```

## Админка:
```
И есть администраторы, которые хотят получать статистику количества запросов пользователей 
по трем вышеперечисленным эндпойнтам. /admin
```

## Запуск проекта:

### Через main.go:
```
1. git clone https://github.com/IKostarev/go-test-inhousead.git
2. убедиться что у вас установлен компилятор Golang, в терминале - go version
3. go run cmd/ping/main.go
4. запуститься проект на порту :8080, при необходимости можно запустить его на другом порту, для этого нужно прописать go run cmd/ping/main.go -addr ":{number}"
5. открыть в браузере localhost:8080/ и затем необходимые ендпоинты
6. в терминале будет дополнительная информация по обработке запросов
```

### Через Docker:
```
1. git clone https://github.com/IKostarev/go-test-inhousead.git
2. в терминале ввести - docker build -t my-golang-app .
3. затем также в терминале - docker run -p 8080:8080 my-golang-app
4. при необходимости изменить порт
```

## Варианты для улучшения приложения:
```
1. Использовать Базу Данных в качестве хранилища.
2. Сделать авторизацию.
```