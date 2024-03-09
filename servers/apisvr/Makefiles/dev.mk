DEV_MYSQL_DSN?='$(shell $(MAKE) -C ../containers/mysql --no-print-directory dsn)'

DEV_SERVER_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)" \
	APP_STAGE=$(APP_STAGE) \
	APP_MYSQL_DSN=$(DEV_MYSQL_DSN) \
	APP_LOG_CONSOLE_WRITER=true \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)

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