include ../config.mk

.PHONY: default
default: build

.PHONY: build
build:
	go build ./...

DBMIGRATIONS_COMMAND_DIR=./bin
$(DBMIGRATIONS_COMMAND_DIR):
	mkdir -p $(DBMIGRATIONS_COMMAND_DIR)

DBMIGRATIONS_COMMAND=$(DBMIGRATIONS_COMMAND_DIR)/runner
$(DBMIGRATIONS_COMMAND): runner_build

.PHONY: runner_build
runner_build: $(DBMIGRATIONS_COMMAND_DIR)
	go build -o $(DBMIGRATIONS_COMMAND) ./cmd/runner && \
	$(DBMIGRATIONS_COMMAND) build-check

DBMIGRATIONS_TABLE?=goose_db_version # goose のデフォルト

DBMIGRATIONS_DSN=$(shell $(MAKE) -C ../containers/mysql --no-print-directory dsn)
.PHONY: dsn
dsn:
	@echo '$(DBMIGRATIONS_DSN)'

# 使用例: DBMIGRATIONS_OPTIONS_EXTRA=--allow-missing make dbmigrate_up
DBMIGRATIONS_OPTIONS_EXTRA?=
DBMIGRATIONS_OPTIONS?=--dir . --table $(DBMIGRATIONS_TABLE) $(DBMIGRATIONS_OPTIONS_EXTRA) mysql '$(DBMIGRATIONS_DSN)'

DBMIGRATIONS_ENVS=\
	STAGE_ENV=$(STAGE_ENV)
.PHONY: envs
envs:
	@echo $(DBMIGRATIONS_ENVS)

# Usage
# make new <migration_name>
ifeq (new,$(firstword $(MAKECMDGOALS)))
  NEW_SQL_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(NEW_SQL_ARGS):;@:)
endif
.PHONY: new
new: $(DBMIGRATIONS_COMMAND)
	$(DBMIGRATIONS_ENVS) $(DBMIGRATIONS_COMMAND) $(DBMIGRATIONS_OPTIONS) create $(NEW_SQL_ARGS) sql

# Usage
# make new_go <migration_name>
ifeq (new_go,$(firstword $(MAKECMDGOALS)))
  NEW_GO_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(NEW_GO_ARGS):;@:)
endif
.PHONY: new_go
new_go: $(DBMIGRATIONS_COMMAND)
	$(DBMIGRATIONS_ENVS) $(DBMIGRATIONS_COMMAND) $(DBMIGRATIONS_OPTIONS) create $(NEW_GO_ARGS) go

# マイグレーション実行
.PHONY: up
up: $(DBMIGRATIONS_COMMAND)
	$(DBMIGRATIONS_ENVS) $(DBMIGRATIONS_COMMAND) $(DBMIGRATIONS_OPTIONS) up

# マイグレーションを一つ戻す
.PHONY: down
down: $(DBMIGRATIONS_COMMAND)
	$(DBMIGRATIONS_ENVS) $(DBMIGRATIONS_COMMAND) $(DBMIGRATIONS_OPTIONS) down

# 状態の確認
.PHONY: state
state: $(DBMIGRATIONS_COMMAND)
	$(DBMIGRATIONS_ENVS) $(DBMIGRATIONS_COMMAND) $(DBMIGRATIONS_OPTIONS) status

# 作成したテーブルを全て削除
.PHONY: reset
reset: $(DBMIGRATIONS_COMMAND)
	$(DBMIGRATIONS_ENVS) $(DBMIGRATIONS_COMMAND) $(DBMIGRATIONS_OPTIONS) reset
