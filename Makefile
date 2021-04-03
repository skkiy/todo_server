.PHONY: local-up
local-up:
	docker-compose -f docker-compose.local.yml up -d --build

.PHONY: local-up-auth
local-up-auth:
	docker-compose -f docker-compose.local.yml build
	docker-compose -f docker-compose.local.yml -v "$(TODO_SECRET_KEY):/secret.json" run

.PHONY: local-up-nocache
local-up-nocache:
	rm -rf docker/mysql/data
	@make local-up

.PHONY: local-down
local-down:
	docker-compose -f docker-compose.local.yml down

.PHONY: generate
generate:
	go run github.com/99designs/gqlgen generate

.PHONY: local-db-exec
local-db-exec:
	docker-compose -f docker-compose.local.yml exec db bash
