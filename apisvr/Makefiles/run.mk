HTTP_PORT?=8000
GRPC_PORT?=8080

APISVR_OPTIONS=\
	-http-port $(HTTP_PORT) \
	-grpc-port $(GRPC_PORT)

SWAGGERUI_PORT?=$(shell $(MAKE) -C ../tools/swaggerui --no-print-directory port)
SWAGGERUI_ORIGIN?="http://localhost:$(SWAGGERUI_PORT)"

APISVR_ENVS=\
	APP_CORS_ALLOW_ORIGINS=$(SWAGGERUI_ORIGIN)

.PHONY: envs
envs:
	@echo $(APISVR_ENVS)

.PHONY: run
run:
	$(APISVR_ENVS) go run ./services/cmd/apisvr $(APISVR_OPTIONS)

.PHONY: run_with_debug
run_with_debug:
	$(APISVR_ENVS) go run ./services/cmd/apisvr $(APISVR_OPTIONS) -debug
