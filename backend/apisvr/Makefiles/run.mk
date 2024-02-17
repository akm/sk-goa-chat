# APP_HTTP_PORT と APP_GRPC_PORT は frontend/tests/integration から呼ばれる際は環境変数で上書きされます
APP_HTTP_PORT?=$(APP_PORT_APISVR_HTTP_dev)
APP_GRPC_PORT?=$(APP_PORT_APISVR_GRPC_dev)

APISVR_OPTIONS=\
	-http-port $(APP_HTTP_PORT) \
	-grpc-port $(APP_GRPC_PORT)

SWAGGERUI_PORT?=$(shell $(MAKE) -C ../../tools/swaggerui --no-print-directory port)
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
