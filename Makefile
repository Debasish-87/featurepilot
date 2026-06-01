include .env
export

run:
	go run cmd/api/main.go

fmt:
	go fmt ./...

test:
	go test ./...

tidy:
	go mod tidy

migrate-up:
	migrate \
	-path migrations \
	-database "$(DB_DSN)" \
	up

migrate-down:
	migrate \
	-path migrations \
	-database "$(DB_DSN)" \
	down 1