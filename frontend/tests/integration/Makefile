include ../../../config.mk
include ./envs.mk

.PHONY: mysql_dsn
mysql_dsn:
	@$(ENVS) $(MAKE) -C ../../../backend/containers/mysql dsn

.PHONY: mysql_wait_to_connect
mysql_wait_to_connect:
	$(ENVS) $(MAKE) -C ../../../backend/containers/mysql wait_to_connect

.PHONY: dbmigrations_up
dbmigrations_up:
	$(ENVS) $(MAKE) -C ../../../backend/dbmigrations up

APISVR_ENVS=\
	GOOGLE_CLOUD_PROJECT=sk-goa-chat \
	APP_HTTP_PORT=$(APP_APISVR_HTTP_PORT) \
	APP_GRPC_PORT=$(APP_APISVR_APP_GRPC_PORT) \
	FIREBASE_AUTH_EMULATOR_HOST=$(FIREBASE_AUTH_EMULATOR_HOST) \

.PHONY: apisvr_envs
apisvr_envs:
	@echo $(APISVR_ENVS)

APISVR_LOG_DIR=log
$(APISVR_LOG_DIR):
	mkdir -p $(APISVR_LOG_DIR)

APISVR_MYSQL_DSN=$(shell $(ENVS) $(MAKE) -C ../../../backend/containers/mysql --no-print-directory dsn)

.PHONY: apisvr_run
apisvr_run: $(APISVR_LOG_DIR)
	$(ENVS) \
	$(APISVR_ENVS) \
	APP_MYSQL_DSN='$(APISVR_MYSQL_DSN)' \
		$(MAKE) -C ../../../backend/apisvr run >> $(APISVR_LOG_DIR)/apisvr.log 2>&1

FIREBASE_AUTH_EMULATOR_HOST=127.0.0.1:$(APP_FIREBASE_AUTH_PORT)

# https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja
#      FIREBASE_AUTH_EMULATOR_HOST と
# VITE_FIREBASE_AUTH_EMULATOR_HOST は重複していますが、どちらの環境変数も必要です
UISVR_ENVS= \
	VITE_APISVR_ORIGIN=http://localhost:$(APP_APISVR_HTTP_PORT) \
	VITE_GOOGLE_CLOUD_PROJECT=sk-goa-chat \
	FIREBASE_AUTH_EMULATOR_HOST=$(FIREBASE_AUTH_EMULATOR_HOST) \
	VITE_FIREBASE_AUTH_EMULATOR_HOST=$(FIREBASE_AUTH_EMULATOR_HOST) \
	VITE_APISVR_GRPC_PORT=$(APP_APISVR_GRPC_PORT)

.PHONY: uisvr_envs
uisvr_envs:
	@echo $(UISVR_ENVS)

.PHONY: uisvr_run
uisvr_run: $(APISVR_LOG_DIR)
	$(ENVS) \
	$(UISVR_ENVS) \
		$(MAKE) -C ../.. build preview >> $(APISVR_LOG_DIR)/uisvr.log 2>&1
