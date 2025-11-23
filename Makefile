DC := docker compose

# DSN для goose с хоста
CONTENT_PG_DSN_LOCAL := postgres://content:secret@localhost:5433/content?sslmode=disable

bin-deps:
	go get github.com/99designs/gqlgen
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/pav5000/smartimports/cmd/smartimports@latest

generate: bin-deps
	@echo "\n --- GraphQL generation --- \n"
	go run github.com/99designs/gqlgen generate

	@echo "\n --- content-service generation --- \n"
	protoc \
      -I api \
      -I $(shell go env GOPATH)/pkg/mod \
      --go_out=internal/pb --go_opt=paths=source_relative \
      --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
      api/content/v1/content.proto

	smartimports .

# Полный перезапуск: стоп → билд → запуск Postgres → миграции → запуск сервиса
.PHONY: run
run:
	@echo "🛑 Останавливаю все контейнеры..."
	$(DC) down --remove-orphans

	@echo "🔧 Пересобираю контейнеры..."
	$(DC) build

	@echo "🐘 Поднимаю Postgres..."
	$(DC) up -d content-postgres

	@echo "⏳ Жду, пока Postgres поднимется..."
	@until $(DC) exec -T content-postgres pg_isready -U content >/dev/null 2>&1 ; do \
		printf "."; \
		sleep 1; \
	done
	@echo "\n✅ Postgres доступен"

	@echo "🐘 Поднимаю pg-exporter..."
	$(DC) up -d postgres-exporter

	@echo "📜 Применяю миграции..."
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(CONTENT_PG_DSN_LOCAL) \
	goose -dir migrations/content-service up
	@echo "✅ Миграции применены"

	@echo "🚀 Поднимаю content-service..."
	$(DC) up -d content-service

	@echo "🌐 Поднимаю gql-api..."
	$(DC) up -d gql-api

	@echo "🧪 Поднимаю prometheus..."
	$(DC) up -d prometheus

	@echo "🧪 Поднимаю grafana..."
	$(DC) up -d grafana

	@echo "✨ Всё поднято! GraphQL: http://localhost:8080/  gRPC: http://localhost:50051  Postgres: http://localhost:5433  Postgres-exporter:http://localhost:9187  Prometheus: http://localhost:9090  Grafana: http://localhost:3000"
