.PHONY: default
default: build

.PHONY: build
build:
	$(MAKE) -C uisvr build

.PHONY: lint
lint:
	$(MAKE) -C uisvr lint

.PHONY: test
test:
	$(MAKE) -C uisvr test_integration_pw_setup && \
	$(MAKE) -C uisvr test

.PHONY: dev
dev:
	$(MAKE) -C uisvr dev
