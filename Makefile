.PHONY: init
init:
	@make up

.PHONY: up
up:
	docker compose up -d

.PHONY: clean
clean:
	docker compose stop
	docker compose down
