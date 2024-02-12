include ../../../.config.mk

ENVS=\
	APP_MYSQL_DATABASE_NAME=$(APP_MYSQL_DATABASE_NAME) \
	APP_MYSQL_PORTT=$(APP_PORT_MYSQL_dev) \
	APP_FIREBASE_AUTH_PORT=$(APP_PORT_FIREBASE_AUTH_dev) \
	APP_FIREBASE_EMULATOR_SUITE_PORT=$(APP_PORT_FIREBASE_EMULATOR_SUITE_dev)

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
