.PHONY: default
default: build

.PHONY: install
install:
	asdf install && \
	asdf reshim

.PHONY: build
build:
	$(MAKE) -C backend build && \
	$(MAKE) -C uisvr build

.PHONY: lint
lint:
	$(MAKE) -C backend lint && \
	$(MAKE) -C uisvr lint

.PHONY: test
test:
	$(MAKE) -C backend test && \
	$(MAKE) -C uisvr test_integration_pw_setup && \
	$(MAKE) -C uisvr test

.PHONY: dev
dev:
	$(MAKE) -C uisvr dev
