ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=homework password=homework dbname=homework host=localhost port=5432 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MIGRATION_FOLDER=$(INTERNAL_PKG_PATH)/db/migrations

up-all:
	docker-compose up -d

down:
	docker-compose down

migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down


unit-tests:
	go test -coverprofile=coverage.out ./internal/pkg/... | grep -v "?"

unit-coverage:
	go tool cover -html=coverage.out

integration-tests:
	go clean -testcache && go test -v ./tests

up-tracing:
	docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
