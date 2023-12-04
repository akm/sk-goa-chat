include ../../config.mk

## mysql クライアントを docker コンテナとして独立して動かす場合はこちらを使う
# LOCAL_CONTAINERS_MYSQL?=localdev_mysql
# LOCAL_CONTAINERS_NETWORK?=sk_goa_chat_containers_localdev_network1
# MYSQL_HOST?=$(shell docker inspect $(LOCAL_CONTAINERS_MYSQL) | jq -r '.[].NetworkSettings.Networks.$(LOCAL_CONTAINERS_NETWORK).IPAddress')
# docker run -it --rm --network=$(LOCAL_CONTAINERS_NETWORK) mysql:8.0.35 mysql -h $(MYSQL_HOST) -P 3306
## asdf でインストールした mysql クライアントを使う場合はこちらを使う
MYSQL_HOST?=127.0.0.1
MYSQL_PORT?=3306
MYSQL_DB_NAME=$(MYSQL_DATABASE_NAME)
MYSQL_USER_NAME?=root
MYSQL_USER_PASSWORD?=

ifeq ($(MYSQL_USER_PASSWORD),)
MYSQL_USER_AUTH_OPTS?=-u $(MYSQL_USER_NAME)
else
MYSQL_USER_AUTH_OPTS?=-u $(MYSQL_USER_NAME) --password=$(MYSQL_USER_PASSWORD)
endif
MYSQL_USER_OPTS=-h $(MYSQL_HOST) -P $(MYSQL_PORT) $(MYSQL_USER_AUTH_OPTS)

.PHONY: wait_to_connect
wait_to_connect:
	MYSQL_USER_OPTS='$(MYSQL_USER_OPTS)' \
	MYSQL_DB_NAME=$(MYSQL_DB_NAME) \
	./wait.sh

.PHONY: console
console:
	mysql $(MYSQL_USER_OPTS) $(MYSQL_DB_NAME)

MYSQL_DSN?=$(MYSQL_USER_NAME):$(MYSQL_USER_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB_NAME)?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo
.PHONY: dsn
dsn:
	@echo "$(MYSQL_DSN)"
