# このファイルは以下の２つのファイルから読み込まれます
# - uisvr/tests/integration/Makefile
# - uisvr/tests/integration/containers/Makefile
ENVS=\
	APP_ENV=e2e_test \
	APP_MYSQL_DATABASE_NAME=$(APP_MYSQL_DATABASE_NAME) \
	APP_MYSQL_PORT=$(APP_PORT_MYSQL_e2e_test) \
	APP_FIREBASE_AUTH_PORT=$(APP_PORT_FIREBASE_AUTH_e2e_test) \
	APP_UISVR_HTTP_PORT=$(APP_PORT_UISVR_HTTP_e2e_test) \
	APP_APISVR_HTTP_PORT=$(APP_PORT_APISVR_HTTP_e2e_test) \
	APP_APISVR_GRPC_PORT=$(APP_PORT_APISVR_GRPC_e2e_test)

.PHONY: envs
envs:
	@echo $(ENVS)
