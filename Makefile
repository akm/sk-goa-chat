.PHONY: default
default: build

.PHONY: install
install:
	asdf install && \
	asdf reshim

.PHONY: build
build:
	$(MAKE) -C servers build && \
	$(MAKE) -C uisvr build

.PHONY: lint
lint:
	$(MAKE) -C servers lint && \
	$(MAKE) -C uisvr lint

.PHONY: test
test:
	$(MAKE) -C servers test && \
	$(MAKE) -C uisvr test_integration_pw_setup && \
	$(MAKE) -C uisvr test

.PHONY: dev
dev:
	$(MAKE) -C uisvr dev
