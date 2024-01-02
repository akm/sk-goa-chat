include ../../../config.mk
include ./envs.mk

.PHONY: mysql_dsn
mysql_dsn:
	@$(ENVS) $(MAKE) -C ../../../containers/mysql dsn

.PHONY: mysql_wait_to_connect
mysql_wait_to_connect:
	$(ENVS) $(MAKE) -C ../../../containers/mysql wait_to_connect

.PHONY: dbmigrations_up
dbmigrations_up:
	$(ENVS) $(MAKE) -C ../../../dbmigrations up

APISVR_ENVS=HTTP_PORT=$(APISVR_HTTP_PORT) GRPC_PORT=$(APISVR_GRPC_PORT)
.PHONY: apisvr_envs
apisvr_envs:
	@echo $(APISVR_ENVS)

APISVR_LOG_DIR=log
$(APISVR_LOG_DIR):
	mkdir -p $(APISVR_LOG_DIR)

APISVR_MYSQL_DSN=$(shell $(ENVS) $(MAKE) -C ../../../containers/mysql --no-print-directory dsn)

.PHONY: apisvr_run
apisvr_run: $(APISVR_LOG_DIR)
	$(ENVS) \
	$(APISVR_ENVS) \
	MYSQL_DSN='$(APISVR_MYSQL_DSN)' \
		$(MAKE) -C ../../../apisvr run >> $(APISVR_LOG_DIR)/apisvr.log 2>&1

FIREBASE_AUTH_EMULATOR_HOST=127.0.0.1:$(FIREBASE_AUTH_PORT)

# https://firebase.google.com/docs/emulator-suite/connect_auth?hl=ja
#      FIREBASE_AUTH_EMULATOR_HOST と
# VITE_FIREBASE_AUTH_EMULATOR_HOST は重複していますが、どちらの環境変数も必要です
UISVR_ENVS= \
	VITE_GOOGLE_CLOUD_PROJECT=sk-goa-chat \
	FIREBASE_AUTH_EMULATOR_HOST=$(FIREBASE_AUTH_EMULATOR_HOST) \
	VITE_FIREBASE_AUTH_EMULATOR_HOST=$(FIREBASE_AUTH_EMULATOR_HOST) \
	VITE_APISVR_GRPC_PORT=$(APISVR_GRPC_PORT)

.PHONY: uisvr_envs
uisvr_envs:
	@echo $(UISVR_ENVS)

.PHONY: uisvr_run
uisvr_run: $(APISVR_LOG_DIR)
	$(ENVS) \
	$(UISVR_ENVS) \
		$(MAKE) -C ../.. build preview >> $(APISVR_LOG_DIR)/uisvr.log 2>&1
