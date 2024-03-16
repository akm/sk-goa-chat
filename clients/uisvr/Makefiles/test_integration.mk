PLAYWRIGHT_ENVS=\
	APP_UISVR_HTTP_PORT=$(APP_PORT_UISVR_HTTP_e2e_test) \
	APP_APISVR_HTTP_PORT=$(APP_PORT_APISVR_HTTP_e2e_test)

.PHONY: test_integration
test_integration: test_integration_reup test_integration_run

.PHONY: test_integration_run
test_integration_run:
	$(PLAYWRIGHT_ENVS) npx playwright test tests/integration/scenarios/*.test.ts

.PHONY: test_integration_reup
test_integration_reup: 
	$(MAKE) -C $(PATH_TO_TEST_INTEGRATION_CONTAINERS) reup

# このターゲットは、playwright のセットアップを行うためのものです。
# 実行すると最新の　playwright のブラウザがインストールされて時間がかかるので、
# test_integration_reup　には含めません。
.PHONY: test_integration_pw_setup
test_integration_pw_setup:
	npx playwright install
