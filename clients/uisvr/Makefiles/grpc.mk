GRPC_SOURCE_DIR=../servers/apisvr/services/gen/grpc
GRPC_DEST_DIR=src/lib/server/protos

GRPC_RESOURCES=users channels chat_messages

.PHONY: grpc_gen
grpc_gen:
	for resource in $(GRPC_RESOURCES); do \
		RESOURCE=$${resource} $(MAKE) grpc_gen_impl; \
	done

.PHONY: grpc_gen_impl
grpc_gen_impl:
	mkdir -p $(GRPC_DEST_DIR) && protoc --ts_out $(GRPC_DEST_DIR) --proto_path $(GRPC_SOURCE_DIR)/$(RESOURCE)/pb $(GRPC_SOURCE_DIR)/$(RESOURCE)/pb/goagen_services_$(RESOURCE).proto
