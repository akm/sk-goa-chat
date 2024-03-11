PATH_TO_PROJECT:=../../..

include $(PATH_TO_PROJECT)/.config.mk

ENVS=\
	APP_MYSQL_DATABASE_NAME=$(APP_MYSQL_DATABASE_NAME) \
	APP_MYSQL_PORT=$(APP_PORT_MYSQL_dev) \
	APP_FIREBASE_AUTH_PORT=$(APP_PORT_FIREBASE_AUTH_dev) \
	APP_FIREBASE_EMULATOR_SUITE_PORT=$(APP_PORT_FIREBASE_EMULATOR_SUITE_dev)

.PHONY: envs
envs:
	@echo $(ENVS)

.PHONY: run
run:
	$(ENVS) docker compose up

.PHONY: up
up: up_start mysql_wait_to_connect dbmigrations_up

.PHONY: mysql_wait_to_connect
mysql_wait_to_connect:
	$(MAKE) -C $(PATH_TO_MYSQL) wait_to_connect

.PHONY: dbmigrations_up
dbmigrations_up:
	$(MAKE) -C $(PATH_TO_DBMIGRATIONS) up

.PHONY: up_start
up_start:
	$(ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS) docker compose down
