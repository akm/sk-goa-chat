# This file uses the following variables or tasks from
#
# Depends on go.mk
#   - GO_ROOT_PACKAGE
GOA_ROOT_PACKAGE=$(GO_ROOT_PACKAGE)
GOA_SERVICES_DIR=./services

GOA=$(shell go env GOPATH)/bin/goa
$(GOA):
	go install goa.design/goa/v3/cmd/goa@v3 && \
	$(MAKE) asdf_reshim

PROTO_GEN_GO=$(shell go env GOPATH)/bin/protoc-gen-go
$(PROTO_GEN_GO):
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
	$(MAKE) asdf_reshim

# protoc は別途インストールする必要がある
PROTO_GEN_GO_GRPC=$(shell go env GOPATH)/bin/protoc-gen-go-grpc
$(PROTO_GEN_GO_GRPC):
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
	$(MAKE) asdf_reshim

GOA_GEN_DIR=./gen
$(GOA_GEN_DIR): goa_gen

.PHONY: goa_gen
goa_gen: $(GOA) $(PROTO_GEN_GO) $(PROTO_GEN_GO_GRPC)
	goa gen $(GOA_ROOT_PACKAGE)/design -o $(GOA_SERVICES_DIR) && \
	$(MAKE) -C ../replacements apisvr_services_gen

.PHONY: services_cmd_remove
services_cmd_remove:
	rm -rf $(GOA_SERVICES_DIR)/cmd

.PHONY: goa_example
goa_example: $(GOA) services_cmd_remove
	goa example $(GOA_ROOT_PACKAGE)/design -o $(GOA_SERVICES_DIR) && \
	$(MAKE) -C ../replacements apisvr_services_cmd
