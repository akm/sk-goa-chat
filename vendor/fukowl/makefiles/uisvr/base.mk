PATH_TO_NODE_MODULES=node_modules
$(PATH_TO_NODE_MODULES):
	$(MAKE) install

.PHONY: install
install:
	npm install

# TODO npm run build ではなく npm run check を使うように変更
.PHONY: build
build: $(PATH_TO_NODE_MODULES) $(BUILD_DEPS)
	npm run build

.PHONY: lint
lint: $(PATH_TO_NODE_MODULES) $(LINT_DEPS)
	npm run lint

.PHONY: preview
preview: $(PATH_TO_NODE_MODULES) $(PREVIEW_DEPS)
	npm run preview

.PHONY: test
test: $(PATH_TO_NODE_MODULES) $(TEST_DEPS)
	npm run test
