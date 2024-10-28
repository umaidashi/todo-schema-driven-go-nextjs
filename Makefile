.PHONY: init
init:
	@make up
	@make logs

.PHONY: up
up:
	docker compose up -d

.PHONY: clean
clean:
	docker compose stop
	docker compose down

.PHONY: gen-ogen
gen-ogen:
	@docker compose run --rm server \
	ogen -package oas -target ./presentation/http/ogen/oas -clean -config ./docs/ogen_config.yaml ./docs/openapi.yaml

.PHONY: gen-ent
gen-ent:
	@docker compose run --rm server \
	go generate ./infra/database/postgres/ent

.PHONY: logs
logs:
	docker compose logs -f
