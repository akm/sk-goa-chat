ENVS_BASE=\
	APP_MYSQL_PORT=$(APP_PORT_MYSQL_unit_test) \
	APP_FIREBASE_AUTH_PORT=$(APP_PORT_FIREBASE_AUTH_unit_test)

.PHONY: envs
envs:
	@echo $(ENVS_BASE) $(ENVS)

.PHONY: run
run:
	$(ENVS_BASE) $(ENVS) docker compose up

.PHONY: up
up: up_start mysql_wait_to_connect dbmigration_up

.PHONY: up_start
up_start:
	$(ENVS_BASE) $(ENVS) docker compose up -d

.PHONY: mysql_wait_to_connect
mysql_wait_to_connect:
	$(ENVS_BASE) $(MAKE) -C $(PATH_TO_MYSQL) wait_to_connect

.PHONY: dbmigration_up
dbmigration_up:
	$(ENVS_BASE) $(MAKE) -C $(PATH_TO_DBMIGRATIONS) up

.PHONY: down
down:
	$(ENVS_BASE) $(ENVS) docker compose down

.PHONY: mysql_dsn
mysql_dsn:
	@$(ENVS_BASE) $(ENVS) $(MAKE) -C $(PATH_TO_MYSQL) dsn
