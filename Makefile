build:
	docker build . -t tender-service

run: build
	docker run -p 8080:8080 tender-service

init-db:
	@echo "==> Initializing database..."
	docker compose exec db psql -U postgres -d tender-service -f /docker-entrypoint-initdb.d/create_table.sql

docker-build:
	@echo "==> Building Docker containers..."
	docker compose build

docker-up: docker-build
	@echo "==> Starting Docker containers..."
	docker compose up

docker-down:
	@echo "==> Stopping Docker containers..."
	docker compose down