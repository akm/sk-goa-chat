ENVS_BASE=\
	APP_MYSQL_PORT=$(APP_PORT_MYSQL_dev) \
	APP_FIREBASE_AUTH_PORT=$(APP_PORT_FIREBASE_AUTH_dev) \
	APP_FIREBASE_EMULATOR_SUITE_PORT=$(APP_PORT_FIREBASE_EMULATOR_SUITE_dev)

.PHONY: envs
envs:
	@echo $(ENVS_BASE) $(ENVS)

.PHONY: run
run:
	$(ENVS_BASE) $(ENVS) docker compose up

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
	$(ENVS_BASE) $(ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS_BASE) $(ENVS) docker compose down
