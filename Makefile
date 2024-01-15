.PHONY: default
default: build

.PHONY: install
install:
	asdf install && \
	asdf reshim

.PHONY: build
build:
	$(MAKE) -C apisvr build && \
	$(MAKE) -C ui build

.PHONY: dev
dev:
	$(MAKE) -C ui dev
