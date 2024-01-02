include ../../../config.mk
include ./envs.mk

.PHONY: mysql_dsn
mysql_dsn:
	@$(ENVS) $(MAKE) -C ../../../containers/mysql dsn

.PHONY: mysql_wait_to_connect
mysql_wait_to_connect:
	$(ENVS) $(MAKE) -C ../../../containers/mysql wait_to_connect

.PHONY: dbmigrations_up
dbmigrations_up:
	$(ENVS) $(MAKE) -C ../../../dbmigrations up

APISVR_ENVS=HTTP_PORT=$(APISVR_HTTP_PORT) GRPC_PORT=$(APISVR_GRPC_PORT)
.PHONY: apisvr_envs
apisvr_envs:
	@echo $(APISVR_ENVS)

APISVR_LOG_DIR=log
$(APISVR_LOG_DIR):
	mkdir -p $(APISVR_LOG_DIR)

APISVR_MYSQL_DSN=$(shell $(ENVS) $(MAKE) -C ../../../containers/mysql --no-print-directory dsn)

.PHONY: apisvr_run
apisvr_run: $(APISVR_LOG_DIR)
	$(ENVS) \
	$(APISVR_ENVS) \
	MYSQL_DSN='$(APISVR_MYSQL_DSN)' \
		$(MAKE) -C ../../../apisvr run >> $(APISVR_LOG_DIR)/apisvr.log 2>&1
