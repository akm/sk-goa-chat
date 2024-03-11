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
up:
	$(ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS) docker compose down

.PHONY: mysql_dsn
mysql_dsn:
	@$(ENVS) $(MAKE) -C $(PATH_TO_MYSQL) dsn
