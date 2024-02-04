.PHONY: default
default: build lint test

.PHONY: build
build:
	$(MAKE) -C applib build && \
	$(MAKE) -C apisvr build

.PHONY: lint
lint:
	$(MAKE) -C applib lint && \
	$(MAKE) -C apisvr lint

.PHONY: test
test:
	$(MAKE) -C applib test && \
	$(MAKE) -C apisvr test

.PHONY: dev
dev:
	$(MAKE) -C apisvr dev
