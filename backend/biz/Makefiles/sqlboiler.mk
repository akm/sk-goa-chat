.PHONY: sqlboiler_gen
sqlboiler_gen: sqlboiler_gen_prepare
	sqlboiler mysql && \
	$(MAKE) -C ../../modifiers backend_biz_models

.PHONY: sqlboiler_gen_prepare
sqlboiler_gen_prepare:
	$(MAKE) -C ../dbmigrations up
