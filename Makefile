.PHONY: init
init:
	@make up
	sleep 5
	@make migrate-up
	@make logs

.PHONY: up
up:
	docker compose up -d

.PHONY: clean
clean:
	docker compose stop
	docker compose down

.PHONY: tidy
tidy:
	@cd server && \
	go mod tidy

.PHONY: gen-ogen
gen-ogen:
	@docker compose run --rm server \
	ogen -package oas -target ./presentation/http/ogen/oas -clean -config ./docs/ogen_config.yaml ./docs/openapi.yaml

.PHONY: gen-ent
gen-ent:
	@docker compose run --rm server \
	go generate ./pkg/ent

.PHONY: create-ddl
create-ddl:
	@docker compose run --rm server \
	go run -mod=mod ./cmd/migration/main.go "$(NAME)"

.PHONY: migrate-up
migrate-up:
	atlas migrate apply \
		--dir "file://server/infra/migrations" \
		--url "postgres://oidon:oidon@localhost:15432/todo-db?sslmode=disable"

.PHONY: logs
logs:
	docker compose logs -f
