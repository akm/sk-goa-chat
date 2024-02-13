TEST_CONTAINER_ENVS=$(shell $(MAKE) -C ../containers/localtest --no-print-directory envs)
TEST_MYSQL_PORT?=$(shell $(TEST_CONTAINER_ENVS) $(MAKE) -C ../containers/mysql --no-print-directory port)
TEST_MYSQL_DSN?='$(shell $(TEST_CONTAINER_ENVS) $(MAKE) -C ../containers/mysql --no-print-directory dsn)'

TEST_ENVS=\
	APP_STAGE=$(APP_STAGE) \
	APP_ENV=unit_test \
	APP_MYSQL_PORT=$(TEST_MYSQL_PORT) \
	APP_MYSQL_DSN=$(TEST_MYSQL_DSN) \
	GOOGLE_CLOUD_PROJECT=sk-goa-chat \
	FIREBASE_AUTH_EMULATOR_HOST='127.0.0.1:9091'

.PHONY: test
test: test_containers_up test_mysql_wait_to_connect test_dbmigration_up
	$(TEST_ENVS) $(MAKE) test_run

.PHONY: test_run
test_run:
	go test -p 1 -tags timetravel ./...

.PHONY: test_containers_up
test_containers_up:
	$(TEST_ENVS) $(MAKE) -C ../containers/localtest up

.PHONY: test_containers_down
test_containers_down:
	$(TEST_ENVS) $(MAKE) -C ../containers/localtest down

.PHONY: test_mysql_wait_to_connect
test_mysql_wait_to_connect:
	$(TEST_ENVS) $(MAKE) -C ../containers/mysql wait_to_connect

.PHONY: test_dbmigration_up
test_dbmigration_up:
	$(TEST_ENVS) $(MAKE) -C ../dbmigrations up
