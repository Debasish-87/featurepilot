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

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-build:
	docker compose build --no-cache

docker-restart:
	docker compose down
	docker compose up -d

logs:
	docker compose logs -f

api-logs:
	docker logs -f featurepilot-api

analytics-logs:
	docker logs -f featurepilot-analytics

db:
	docker exec -it featurepilot-postgres psql -U postgres -d featurepilot

status:
	docker ps

clean:
	docker compose down
	docker system prune -f