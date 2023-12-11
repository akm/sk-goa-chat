DEV_MYSQL_DSN?='$(shell $(MAKE) -C ../containers/mysql --no-print-directory dsn)'

DEV_SERVER_ENVS=\
	STAGE=$(STAGE) \
	STAGE_ENV=$(STAGE_ENV) \
	MYSQL_DSN=$(DEV_MYSQL_DSN)

.PHONY: dev
dev: dev_containers_up
	$(DEV_SERVER_ENVS) $(MAKE) dev_run

.PHONY: dev_run
dev_run:
	go run ./services/cmd/apisvr

.PHONY: dev_containers_up
dev_containers_up:
	$(DEV_SERVER_ENVS) $(MAKE) -C ../containers/localdev up

.PHONY: dev_containers_down
dev_containers_down:
	$(DEV_SERVER_ENVS) $(MAKE) -C ../containers/localdev down
