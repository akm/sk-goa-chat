.PHONY: default
default: build

.PHONY: install
install:
	asdf install && \
	asdf reshim

.PHONY: build
build:
	$(MAKE) -C backend build && \
	$(MAKE) -C frontend build

.PHONY: lint
lint:
	$(MAKE) -C backend lint && \
	$(MAKE) -C frontend lint

.PHONY: test
test:
	$(MAKE) -C backend test && \
	$(MAKE) -C frontend test_integration_pw_setup && \
	$(MAKE) -C frontend test

.PHONY: dev
dev:
	$(MAKE) -C frontend dev
