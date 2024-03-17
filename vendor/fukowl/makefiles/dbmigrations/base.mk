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

DSN=$(shell $(MAKE) -C $(PATH_TO_MYSQL) --no-print-directory dsn)
.PHONY: dsn
dsn:
	@echo '$(DSN)'

# 使用例: OPTIONS_EXTRA=--allow-missing make dbmigrate_up
OPTIONS_EXTRA?=
OPTIONS?=--dir . --table $(TABLE) $(OPTIONS_EXTRA) mysql '$(DSN)'

ENVS=\
	APP_STAGE_TYPE=$(APP_STAGE_TYPE) \
	APP_STAGE=$(APP_STAGE) \
	APP_ENV=$(APP_ENV)
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
	$(ENVS) $(MAKE) dump_schema

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

# 同じディレクトリの .sql ファイルはマイグレーションとしてみなされてしまうので、親ディレクトリに置きます
SCHEMA_DUMP_SQL_FILE?=../schema.dump.sql
$(SCHEMA_DUMP_SQL_FILE): dump_schema

MYSQLDUMP_OPTIONS=$(shell $(MAKE) -C ../containers/mysql --no-print-directory cli_options)

.PHONY: dump_schema
dump_schema:
	echo "APP_SKIP_DB_SCHEMA_DUMP: $$APP_SKIP_DB_SCHEMA_DUMP, APP_ENV: $$APP_ENV"
	( [ -n "$$APP_SKIP_DB_SCHEMA_DUMP" ] || [[ $$APP_ENV =~ test ]] ) && \
		echo "SKIP dump schema" || \
		mysqldump $(MYSQLDUMP_OPTIONS) --no-data $(APP_MYSQL_DATABASE_NAME) > $(SCHEMA_DUMP_SQL_FILE)
