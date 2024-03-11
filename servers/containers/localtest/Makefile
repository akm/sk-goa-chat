PATH_TO_PROJECT:=../../..

include $(PATH_TO_PROJECT)/.config.mk

ENVS=\
	APP_MYSQL_DATABASE_NAME=$(APP_MYSQL_DATABASE_NAME) \
	APP_MYSQL_PORT=$(APP_PORT_MYSQL_unit_test) \
	APP_FIREBASE_AUTH_PORT=$(APP_PORT_FIREBASE_AUTH_unit_test)

.PHONY: envs
envs:
	@echo $(ENVS)

.PHONY: run
run:
	$(ENVS) docker compose up

.PHONY: up
up: up_start mysql_wait_to_connect dbmigration_up

.PHONY: up_start
up_start:
	$(ENVS) docker compose up -d

.PHONY: mysql_wait_to_connect
mysql_wait_to_connect:
	$(MAKE) -C $(PATH_TO_MYSQL) wait_to_connect

.PHONY: dbmigration_up
dbmigration_up:
	$(MAKE) -C $(PATH_TO_DBMIGRATIONS) up

.PHONY: down
down:
	$(ENVS) docker compose down

.PHONY: mysql_dsn
mysql_dsn:
	@$(ENVS) $(MAKE) -C $(PATH_TO_MYSQL) dsn
