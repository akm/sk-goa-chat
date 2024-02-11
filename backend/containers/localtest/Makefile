include ../../../config.mk

ENVS=\
	APP_MYSQL_DATABASE_NAME=$(APP_MYSQL_DATABASE_NAME) \
	APP_MYSQL_PORT=3311

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

.PHONY: mysql_dsn
mysql_dsn:
	@$(ENVS) $(MAKE) -C ../mysql dsn
