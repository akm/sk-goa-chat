.PHONY: grpc_gen
grpc_gen:
	for resource in $(GRPC_RESOURCES); do \
		RESOURCE=$${resource} $(MAKE) grpc_gen_impl; \
	done

.PHONY: grpc_gen_impl
grpc_gen_impl:
	mkdir -p $(PATH_TO_GRPC_DEST_DIR) && \
	npx protoc \
		--ts_out $(PATH_TO_GRPC_DEST_DIR) \
		--proto_path $(PATH_TO_GRPC_SOURCE_DIR)/$(RESOURCE)/pb \
		$(PATH_TO_GRPC_SOURCE_DIR)/$(RESOURCE)/pb/goagen_services_$(RESOURCE).proto
