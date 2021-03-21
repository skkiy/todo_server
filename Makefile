.PHONY: local-up
local-up:
	docker-compose -f docker-compose.local.yml up -d --build

.PHONY: local-up-nocache
local-up-nocache:
	rm -rf docker/mysql/data
	@make local-up

.PHONY: local-down
local-down:
	docker-compose -f docker-compose.local.yml down
