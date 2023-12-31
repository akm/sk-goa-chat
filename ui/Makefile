include ../config.mk
include ./envs.mk

.PHONY: mysql_dsn
mysql_dsn:
	@$(ENVS) $(MAKE) -C ../containers/mysql dsn

.PHONY: dbmigrations_up
dbmigrations_up:
	$(ENVS) $(MAKE) -C ../dbmigrations up

APISVR_ENVS=HTTP_PORT=$(APISVR_HTTP_PORT) GRPC_PORT=$(APISVR_GRPC_PORT)
.PHONY: apisvr_envs
apisvr_envs:
	@echo $(APISVR_ENVS)

.PHONY: apisvr_run
apisvr_run:
	$(ENVS) $(APISVR_ENVS) $(MAKE) -C ../apisvr run
