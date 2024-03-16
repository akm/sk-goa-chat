CONTAINERS_ENVS=\
	PATH_TO_PROJECT=$(PATH_TO_PROJECT)

.PHONY: run
run:
	$(ENVS) $(CONTAINERS_ENVS) docker compose up

.PHONY: reup
reup: down up

.PHONY: restart
restart: down up_start

.PHONY: up
up: up_start mysql_wait_to_connect dbmigrations_up

.PHONY: up_start
up_start:
	$(ENVS) $(CONTAINERS_ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS) $(CONTAINERS_ENVS) docker compose down

.PHONY: mysql_wait_to_connect
mysql_wait_to_connect:
	$(ENVS) $(MAKE) -C $(PATH_TO_MYSQL) wait_to_connect

.PHONY: dbmigrations_up
dbmigrations_up:
	$(ENVS) $(MAKE) -C $(PATH_TO_DBMIGRATIONS) up
