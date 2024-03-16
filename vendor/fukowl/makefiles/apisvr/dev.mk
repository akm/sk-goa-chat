DEV_MYSQL_DSN?='$(shell $(MAKE) -C $(PATH_TO_MYSQL) --no-print-directory dsn)'

DEV_SERVER_ENVS_BASE=\
	APP_STAGE=$(APP_STAGE) \
	APP_MYSQL_DSN=$(DEV_MYSQL_DSN) \
	APP_LOG_CONSOLE_WRITER=true \
	FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)"

.PHONY: dev
dev: dev_containers_up
	$(DEV_SERVER_ENVS_BASE) $(DEV_SERVER_ENVS) $(MAKE) run_with_debug

.PHONY: dev_containers_up
dev_containers_up:
	$(DEV_SERVER_ENVS_BASE) $(DEV_SERVER_ENVS) $(MAKE) -C $(PATH_TO_LOCALDEV) up

.PHONY: dev_containers_down
dev_containers_down:
	$(DEV_SERVER_ENVS_BASE) $(DEV_SERVER_ENVS) $(MAKE) -C $(PATH_TO_LOCALDEV) down
