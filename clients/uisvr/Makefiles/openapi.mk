OPENAPI_TS_FILE=src/lib/openapi.d.ts
$(OPENAPI_TS_FILE):
	$(MAKE) openapi_gen

.PHONY: openapi_gen
openapi_gen:
	npx openapi-typescript $(PATH_TO_PROJECT)/servers/apisvr/services/gen/http/openapi3.yaml -o $(OPENAPI_TS_FILE)
