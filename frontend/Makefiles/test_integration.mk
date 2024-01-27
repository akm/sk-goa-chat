.PHONY: test_integration
test_integration:
	npm run test:integration

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
