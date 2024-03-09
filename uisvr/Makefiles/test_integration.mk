PLAYWRIGHT_ENVS=\
	APP_UISVR_HTTP_PORT=$(APP_PORT_UISVR_HTTP_e2e_test) \
	APP_APISVR_HTTP_PORT=$(APP_PORT_APISVR_HTTP_e2e_test)

.PHONY: test_integration
test_integration: test_integration_setup test_integration_run

.PHONY: test_integration_run
test_integration_run: test_integration_setup
	$(PLAYWRIGHT_ENVS) npx playwright test tests/integration/scenarios/*.test.ts

.PHONY: test_integration_setup
test_integration_setup: test_integration_containers_restart test_integration_dbmigrations_up

.PHONY: test_integration_containers_up
test_integration_containers_up:
	$(MAKE) -C tests/integration/containers up

.PHONY: test_integration_containers_down
test_integration_containers_down:
	$(MAKE) -C tests/integration/containers down

.PHONY: test_integration_containers_restart
test_integration_containers_restart:
	$(MAKE) -C tests/integration/containers down up

.PHONY: test_integration_dbmigrations_up
test_integration_dbmigrations_up:
	$(MAKE) -C tests/integration mysql_wait_to_connect dbmigrations_up

.PHONY: test_integration_setup
test_integration_setup:
	$(MAKE) -C tests/integration test_setup

# このターゲットは、playwright のセットアップを行うためのものです。
# 実行すると最新の　playwright のブラウザがインストールされて時間がかかるので、
# test_integration_setup　には含めません。
.PHONY: test_integration_pw_setup
test_integration_pw_setup:
	npx playwright install
