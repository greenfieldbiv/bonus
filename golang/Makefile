# Запуск базы данных (пока не работает)
db_up:
	docker-compose up -d

# Сборка бинарника (пока не работает)
build:
	go build -o ./tmp/go-app/main cmd/main/main.go

# Запуск приложения (пока не работает)
run:
	go run cmd/main/main.go

# Сборка контейнера сервиса (пока не работает)
docker_app_build: build
	docker build ./deploy

# Сваггер (пока не работает)
swaginit:
	swag init
