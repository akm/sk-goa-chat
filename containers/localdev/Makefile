include ../../config.mk

ENVS=\
	MYSQL_DATABASE_NAME=$(MYSQL_DATABASE_NAME)

.PHONY: envs
envs:
	@echo $(ENVS)

.PHONY: run
run:
	$(ENVS) docker compose up

.PHONY: up
up:
	$(ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS) docker compose down
