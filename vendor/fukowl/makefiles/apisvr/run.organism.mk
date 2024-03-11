# APP_HTTP_PORT と APP_GRPC_PORT は clients/uisvr/tests/integration から呼ばれる際は環境変数で上書きされます
APP_HTTP_PORT?=$(APP_PORT_APISVR_HTTP_dev)
APP_GRPC_PORT?=$(APP_PORT_APISVR_GRPC_dev)

APISVR_OPTIONS=\
	-http-port $(APP_HTTP_PORT) \
	-grpc-port $(APP_GRPC_PORT)

SWAGGERUI_PORT?=$(shell $(MAKE) -C $(PATH_TO_SWAGGERUI) --no-print-directory port)
SWAGGERUI_ORIGIN?="http://localhost:$(SWAGGERUI_PORT)"

APISVR_ENVS=\
	APP_CORS_ALLOW_ORIGINS=$(SWAGGERUI_ORIGIN)

.PHONY: apisvr_envs
apisvr_envs:
	@echo $(APISVR_ENVS)

.PHONY: run
run:
	$(APISVR_ENVS) go run $(PATH_TO_CMD_APISVR) $(APISVR_OPTIONS)

.PHONY: run_with_debug
run_with_debug:
	APISVR_OPTIONS='$(APISVR_OPTIONS) -debug' $(MAKE) run
