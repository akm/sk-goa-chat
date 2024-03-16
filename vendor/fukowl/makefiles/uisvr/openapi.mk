$(PATH_TO_OPENAPI_TS_FILE):
	$(MAKE) openapi_gen

.PHONY: openapi_gen
openapi_gen:
	npx openapi-typescript $(PATH_TO_OPENAPI_DEFINITION_FILE) -o $(PATH_TO_OPENAPI_TS_FILE)
