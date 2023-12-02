# https://github.com/golangci/golangci-lint
LINT_GOLANGCI_LINT_VERSION=v1.55.2
LINT_GOLANGCI_LINT=$(shell go env GOPATH)/bin/golangci-lint
$(LINT_GOLANGCI_LINT):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINT_GOLANGCI_LINT_VERSION)

.PHONY: lint
lint: $(LINT_GOLANGCI_LINT)
	$(LINT_GOLANGCI_LINT) run ./...
