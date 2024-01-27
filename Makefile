.PHONY: default
default: build

.PHONY: install
install:
	asdf install && \
	asdf reshim

.PHONY: build
build:
	$(MAKE) -C apisvr build && \
	$(MAKE) -C frontend build

.PHONY: test
test:
	$(MAKE) -C apisvr test && \
	$(MAKE) -C frontend test

.PHONY: dev
dev:
	$(MAKE) -C frontend dev
