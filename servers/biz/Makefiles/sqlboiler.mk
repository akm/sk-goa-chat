SQLBOILER=$(shell go env GOPATH)/bin/sqlboiler
$(SQLBOILER):
	go install github.com/volatiletech/sqlboiler/v4@latest && \
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest && \
	asdf reshim

.PHONY: sqlboiler_gen
sqlboiler_gen: $(SQLBOILER) sqlboiler_gen_prepare
	sqlboiler mysql && \
	$(MAKE) -C ../replacements biz_models

.PHONY: sqlboiler_gen_ci
sqlboiler_gen_ci: sqlboiler_gen git_check_uncommited_changes git_check_untracked_files

.PHONY: sqlboiler_gen_prepare
sqlboiler_gen_prepare:
	$(MAKE) -C ../containers/localdev up && \
	$(MAKE) -C ../containers/mysql wait_to_connect && \
	APP_SKIP_DB_SCHEMA_DUMP=true $(MAKE) -C ../dbmigrations up
