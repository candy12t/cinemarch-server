BIN := bin/cs

ENV_LOCAL_FILE := ./env.local
ENV_LOCAL = $(shell cat $(ENV_LOCAL_FILE))

.PHONY: build
build: test clean
	go build -o $(BIN) -v ./cmd/cinema-search

.PHONY: serve
serve:
	$(ENV_LOCAL) go run ./cmd/cinema-search/main.go

.PHONY: test
test:
	$(ENV_LOCAL) go test ./... -count=1 --race -v

.PHONY: clean
clean:
	rm -rf $(BIN)
	go clean

.PHONY: run-local-db
run-local-db:
	set -a && \
		. $(ENV_LOCAL_FILE) && \
		docker compose up -d && \
		until mysql -u$${DB_USER} -p$${DB_PASSWORD} -h$${DB_HOST} -P$${DB_PORT} -e "SELECT 1"; do sleep 10; done

.PHONY: migrate
migrate:
	set -a && \
		. $(ENV_LOCAL_FILE) && \
		mysql -u$${DB_USER} -p$${DB_PASSWORD} -h$${DB_HOST} -P$${DB_PORT} < _tools/mysql/schema.sql

.PHONY: stop-local-db
stop-local-db:
	docker compose down

.PHONY: delete-local-db
delete-local-db:
	docker compose down -v
