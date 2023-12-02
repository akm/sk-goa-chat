GO_ROOT_PACKAGE=apisvr

.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test ./...
