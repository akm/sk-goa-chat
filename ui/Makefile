include ../config.mk
include ./envs.mk

.PHONY: mysql_dsn
mysql_dsn:
	@$(ENVS) $(MAKE) -C ../containers/mysql dsn

.PHONY: dbmigrations_up
dbmigrations_up:
	$(ENVS) $(MAKE) -C ../dbmigrations up

.PHONY: apisvr_run
apisvr_run:
	$(ENVS) $(MAKE) -C ../apisvr run
