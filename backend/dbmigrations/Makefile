include ../../.config.mk

.PHONY: default
default: build

.PHONY: build
build:
	go build ./...

COMMAND_DIR=./bin
$(COMMAND_DIR):
	mkdir -p $(COMMAND_DIR)

COMMAND_RUNNER=$(COMMAND_DIR)/runner
$(COMMAND_RUNNER): runner_build

.PHONY: runner_build
runner_build: $(COMMAND_DIR)
	go build -o $(COMMAND_RUNNER) ./cmd/runner && \
	$(COMMAND_RUNNER) build-check

TABLE?=goose_db_version # goose のデフォルト

DSN=$(shell $(MAKE) -C ../containers/mysql --no-print-directory dsn)
.PHONY: dsn
dsn:
	@echo '$(DSN)'

# 使用例: OPTIONS_EXTRA=--allow-missing make dbmigrate_up
OPTIONS_EXTRA?=
OPTIONS?=--dir . --table $(TABLE) $(OPTIONS_EXTRA) mysql '$(DSN)'

ENVS=
.PHONY: envs
envs:
	@echo $(ENVS)

# Usage
# make new <migration_name>
ifeq (new,$(firstword $(MAKECMDGOALS)))
  NEW_SQL_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(NEW_SQL_ARGS):;@:)
endif
.PHONY: new
new: $(COMMAND_RUNNER)
	$(ENVS) $(COMMAND_RUNNER) $(OPTIONS) create $(NEW_SQL_ARGS) sql

# Usage
# make new_go <migration_name>
ifeq (new_go,$(firstword $(MAKECMDGOALS)))
  NEW_GO_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(NEW_GO_ARGS):;@:)
endif
.PHONY: new_go
new_go: $(COMMAND_RUNNER)
	$(ENVS) $(COMMAND_RUNNER) $(OPTIONS) create $(NEW_GO_ARGS) go

# マイグレーション実行
.PHONY: up
up: $(COMMAND_RUNNER)
	$(ENVS) $(COMMAND_RUNNER) $(OPTIONS) up && \
	$(MAKE) dump_schema

# マイグレーションを一つ戻す
.PHONY: down
down: $(COMMAND_RUNNER)
	$(ENVS) $(COMMAND_RUNNER) $(OPTIONS) down

# 状態の確認
.PHONY: state
state: $(COMMAND_RUNNER)
	$(ENVS) $(COMMAND_RUNNER) $(OPTIONS) status

# 作成したテーブルを全て削除
.PHONY: reset
reset: $(COMMAND_RUNNER)
	$(ENVS) $(COMMAND_RUNNER) $(OPTIONS) reset

SCHEMA_DUMP_SQL_FILE?=schema.dump.sql
$(SCHEMA_DUMP_SQL_FILE): dump_schema

MYSQLDUMP_OPTIONS=$(shell $(MAKE) -C ../containers/mysql cli_options)

.PHONY: dump_schema
dump_schema:
	mysqldump $(MYSQLDUMP_OPTIONS) --no-data $(APP_MYSQL_DATABASE_NAME) > $(SCHEMA_DUMP_SQL_FILE)