# Сервис для создания и просмотра новостей

## Установка и запуск

1. Клонируем репозиторий

2. Выполняем команду, предварительно настроив .env

```sh
# Для запуска через docker-compose

docker-compose up
```

### Как установить Protocol Buffer на Windows (опционально)
1. Скачать Protocol Buffer для своей операционной системы. [Ссылка Github](https://github.com/protocolbuffers/protobuf/releases)
2. Распаковать архив и переместить файл с расширением .bin на диск C
3. Скопировать путь и задать его в переменной окружения (Система -> Изменение системных переменных среды -> Переменные среды... -> Два раза кликнуть на PATH -> Создать -> Вставить скопированный путь -> ОК)
4. Установить `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
5. Установить `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
6. Скачать `go get -u github.com/golang/protobuf/protoc-gen-go`
7. Проверить в терминале командой `protoc`, должен отобразиться текст с информацией о прото.