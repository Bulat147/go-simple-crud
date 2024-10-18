include .env # Это файл для переменных окружения его нужно иметь у себя локально с теми переменными, что в этом файле используются

build:
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux ./cmd/app/main.go
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows ./cmd/app/main.go

migrate-create:
	migrate create -ext sql -dir db/migrations -seq create_tasks_table

migrate-up:
	migrate -path db/migrations -database "${DB_URL}" up

migrate-down: # здесь yes - чтобы автоматически отвечать да на все вопросы
	yes | migrate -path db/migrations -database "${DB_URL}" down

compose-up: # тут флаг -d обозначает запуск в фоновом режиме, чтобы он не перекрывал cmd своей работой
	docker-compose --project-directory docker --env-file .env up -d 

compose-down:
	docker-compose --project-directory docker down

sleep: # Между командами, иначе compose-up и migrate-up ломаются совместно
	sleep 8

run: build compose-up sleep migrate-up
	./target/${BINARY_NAME}-${CURRENT_OS}

clean: migrate-down compose-down
	go clean
	rm ./target/${BINARY_NAME}-linux
	rm ./target/${BINARY_NAME}-windows

