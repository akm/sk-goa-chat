# TODO sqlboiler の接続先を localdev ではなく localtest に変更

SQLBOILER=$(shell go env GOPATH)/bin/sqlboiler
$(SQLBOILER):
	go install github.com/volatiletech/sqlboiler/v4@latest && \
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest && \
	$(MAKE) asdf_reshim

.PHONY: sqlboiler_gen
sqlboiler_gen: $(SQLBOILER) sqlboiler_gen_prepare
	sqlboiler mysql && \
	$(MAKE) -C $(PATH_TO_SERVER_REPLACEMENTS) biz_models

.PHONY: sqlboiler_gen_prepare
sqlboiler_gen_prepare:
	$(MAKE) -C $(PATH_TO_LOCALDEV) up && \
	$(MAKE) -C $(PATH_TO_MYSQL) wait_to_connect && \
	APP_SKIP_DB_SCHEMA_DUMP=true $(MAKE) -C $(PATH_TO_DBMIGRATIONS) up
