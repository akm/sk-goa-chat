## mysql クライアントを docker コンテナとして独立して動かす場合はこちらを使う
# LOCAL_CONTAINERS_MYSQL?=localdev_mysql
# LOCAL_CONTAINERS_NETWORK?=sk_goa_chat_containers_localdev_network1
# APP_MYSQL_HOST?=$(shell docker inspect $(LOCAL_CONTAINERS_MYSQL) | jq -r '.[].NetworkSettings.Networks.$(LOCAL_CONTAINERS_NETWORK).IPAddress')
# docker run -it --rm --network=$(LOCAL_CONTAINERS_NETWORK) mysql:8.0.35 mysql -h $(APP_MYSQL_HOST) -P 3306
## asdf でインストールした mysql クライアントを使う場合はこちらを使う
APP_MYSQL_HOST?=127.0.0.1
APP_MYSQL_PORT?=3306
APP_MYSQL_USER_NAME?=root
APP_MYSQL_USER_PASSWORD?=

ifeq ($(APP_MYSQL_USER_PASSWORD),)
APP_MYSQL_USER_AUTH_OPTS?=-u $(APP_MYSQL_USER_NAME)
else
APP_MYSQL_USER_AUTH_OPTS?=-u $(APP_MYSQL_USER_NAME) --password=$(APP_MYSQL_USER_PASSWORD)
endif
APP_MYSQL_USER_OPTS=-h $(APP_MYSQL_HOST) -P $(APP_MYSQL_PORT) $(APP_MYSQL_USER_AUTH_OPTS)

.PHONY: cli_options
cli_options:
	@echo $(APP_MYSQL_USER_OPTS)

.PHONY: wait_to_connect
wait_to_connect:
	APP_MYSQL_USER_OPTS='$(APP_MYSQL_USER_OPTS)' \
	APP_MYSQL_DB_NAME=$(APP_MYSQL_DB_NAME) \
	./wait.sh

.PHONY: console
console:
	mysql $(APP_MYSQL_USER_OPTS) $(APP_MYSQL_DB_NAME)

APP_MYSQL_DSN?=$(APP_MYSQL_USER_NAME):$(APP_MYSQL_USER_PASSWORD)@tcp($(APP_MYSQL_HOST):$(APP_MYSQL_PORT))/$(APP_MYSQL_DB_NAME)?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo
.PHONY: dsn
dsn:
	@echo "$(APP_MYSQL_DSN)"

.PHONY: port
port:
	@echo "$(APP_MYSQL_PORT)"
