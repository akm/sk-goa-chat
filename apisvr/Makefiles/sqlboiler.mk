.PHONY: sqlboiler_gen
sqlboiler_gen: sqlboiler_gen_prepare
	sqlboiler mysql

.PHONY: sqlboiler_gen_prepare
sqlboiler_gen_prepare:
	$(MAKE) -C ../dbmigrations up
