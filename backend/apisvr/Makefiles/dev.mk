DEV_MYSQL_DSN?='$(shell $(MAKE) -C ../containers/mysql --no-print-directory dsn)'

DEV_SERVER_ENVS=\
	APP_STAGE=$(APP_STAGE) \
	MYSQL_DSN=$(DEV_MYSQL_DSN) \
	LOG_CONSOLE_WRITER=true

.PHONY: dev
dev: dev_containers_up
	$(DEV_SERVER_ENVS) $(MAKE) run_with_debug

.PHONY: dev_containers_up
dev_containers_up:
	$(DEV_SERVER_ENVS) $(MAKE) -C ../containers/localdev up

.PHONY: dev_containers_up_with_migration
dev_containers_up_with_migration: dev_containers_up dev_dbmigrations_up

.PHONY: dev_containers_down
dev_containers_down:
	$(DEV_SERVER_ENVS) $(MAKE) -C ../containers/localdev down

.PHONY: dev_mysql_wait_to_connect
dev_mysql_wait_to_connect:
	$(ENVS) $(MAKE) -C ../containers/mysql wait_to_connect

.PHONY: dev_dbmigrations_up
dev_dbmigrations_up: dev_mysql_wait_to_connect
	$(ENVS) $(MAKE) -C ../dbmigrations up
