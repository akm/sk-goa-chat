.PHONY: sqlboiler_gen
sqlboiler_gen: sqlboiler_gen_prepare
	sqlboiler mysql && \
	$(MAKE) -C ../../modifiers backend_biz_models

.PHONY: sqlboiler_gen_prepare
sqlboiler_gen_prepare:
	$(MAKE) -C ../containers/localdev up && \
	$(MAKE) -C ../containers/mysql wait_to_connect && \
	APP_SKIP_DB_SCHEMA_DUMP=true $(MAKE) -C ../dbmigrations up
