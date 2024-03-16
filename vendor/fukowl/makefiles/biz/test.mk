TEST_CONTAINER_ENVS=$(shell $(MAKE) -C $(PATH_TO_LOCALTEST) --no-print-directory envs)
TEST_MYSQL_PORT?=$(shell $(TEST_CONTAINER_ENVS) $(MAKE) -C $(PATH_TO_MYSQL) --no-print-directory port)
TEST_MYSQL_DSN?='$(shell $(TEST_CONTAINER_ENVS) $(MAKE) -C $(PATH_TO_MYSQL) --no-print-directory dsn)'

TEST_ENVS=\
	APP_STAGE=$(APP_STAGE) \
	APP_ENV=unit_test \
	APP_MYSQL_PORT=$(TEST_MYSQL_PORT) \
	APP_MYSQL_DSN=$(TEST_MYSQL_DSN) \
	MYSQL_PORT=$(TEST_MYSQL_PORT) \
	FIREBASE_AUTH_EMULATOR_HOST='127.0.0.1:$(APP_PORT_FIREBASE_AUTH_unit_test)'

# 環境変数 MYSQL_PORT はどこにも使われていないように見えますが
# servers/biz/models/mysql_main_test.go で設定される
# mysql.port に対応しています。

.PHONY: test
test: test_containers_up
	$(TEST_ENVS) $(MAKE) test_run

.PHONY: test_run
test_run:
	go test -p 1 -tags timetravel ./...

.PHONY: test_containers_up
test_containers_up:
	$(TEST_ENVS) $(MAKE) -C $(PATH_TO_LOCALTEST) up

.PHONY: test_containers_down
test_containers_down:
	$(TEST_ENVS) $(MAKE) -C $(PATH_TO_LOCALTEST) down
