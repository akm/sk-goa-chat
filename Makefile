.PHONY: default
default: build

.PHONY: install
install:
	asdf install && \
	asdf reshim

.PHONY: build
build:
	$(MAKE) -C servers build && \
	$(MAKE) -C clients build

.PHONY: lint
lint:
	$(MAKE) -C servers lint && \
	$(MAKE) -C clients lint

.PHONY: test
test:
	$(MAKE) -C servers test && \
	$(MAKE) -C clients test

.PHONY: dev
dev:
	$(MAKE) -C clients dev
