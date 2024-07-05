.PHONY: build
db:
	docker-compose up -d --build

migrateup:
	migrate -path migrations -database "postgres://root:root@localhost:5444/time_tracker?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgres://root:root@localhost:5444/time_tracker?sslmode=disable" -verbose down -1

http:
	go run ./cmd/main.go

build: db migrateup http


.DEFAULT_GOAL := build