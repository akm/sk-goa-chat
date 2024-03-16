# https://github.com/golangci/golangci-lint
# See https://golangci-lint.run/usage/false-positives/#nolint-directive for nolint directive
ifndef LINT_GOLANGCI_LINT_VERSION
LINT_GOLANGCI_LINT_VERSION=v1.55.2
endif
LINT_GOLANGCI_LINT=$(shell go env GOPATH)/bin/golangci-lint
$(LINT_GOLANGCI_LINT):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINT_GOLANGCI_LINT_VERSION) && \
	$(MAKE) asdf_reshim

.PHONY: lint
lint: $(LINT_GOLANGCI_LINT)
	$(LINT_GOLANGCI_LINT) run ./...
