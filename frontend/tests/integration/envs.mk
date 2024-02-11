# このファイルは以下の２つのファイルから読み込まれます
# - frontend/tests/integration/Makefile
# - frontend/tests/integration/containers/Makefile
APISVR_HTTP_PORT=8001
APISVR_GRPC_PORT=8081
FIREBASE_AUTH_PORT=9090
ENVS=\
	APP_ENV=e2e_test \
	APP_MYSQL_DATABASE_NAME=$(APP_MYSQL_DATABASE_NAME) \
	APP_MYSQL_PORT=3307 \
	FIREBASE_AUTH_PORT=$(FIREBASE_AUTH_PORT) \
	APISVR_HTTP_PORT=$(APISVR_HTTP_PORT) \
	APISVR_GRPC_PORT=$(APISVR_GRPC_PORT)

.PHONY: envs
envs:
	@echo $(ENVS)
