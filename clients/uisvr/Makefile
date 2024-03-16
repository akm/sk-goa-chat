PATH_TO_PROJECT:=../..
PATH_TO_FUKOWL_ROOT=$(PATH_TO_PROJECT)/vendor/fukowl
PATH_TO_FUKOWL_MAKEFILES=$(PATH_TO_FUKOWL_ROOT)/makefiles

include $(PATH_TO_PROJECT)/.config.mk

.PHONY: default
default: build lint test

PATH_TO_FIREBASE_CONFIG_FILE=src/lib/firebase/firebaseconfig.ts
$(PATH_TO_FIREBASE_CONFIG_FILE):
	cp $(PATH_TO_FIREBASE_CONFIG_FILE).dummy $(PATH_TO_FIREBASE_CONFIG_FILE)

PATH_TO_NODE_MODULES=node_modules
$(PATH_TO_NODE_MODULES):
	$(MAKE) install

BASE_DEPENDENCIES=$(PATH_TO_NODE_MODULES) $(PATH_TO_FIREBASE_CONFIG_FILE)
BUILD_DEPENDENCIES   = $(BASE_DEPENDENCIES)
LINT_DEPENDENCIES    = $(BASE_DEPENDENCIES)
TEST_DEPENDENCIES    = $(BASE_DEPENDENCIES)
PREVIEW_DEPENDENCIES = $(BASE_DEPENDENCIES)
DEV_DEPENDENCIES     = $(BASE_DEPENDENCIES)

.PHONY: install
install:
	npm install

# TODO npm run build ではなく npm run check を使うように変更
.PHONY: build
build: $(BUILD_DEPENDENCIES)
	npm run build

.PHONY: lint
lint: $(LINT_DEPENDENCIES)
	npm run lint

.PHONY: preview
preview: $(PREVIEW_DEPENDENCIES)
	npm run preview

.PHONY: test
test: $(TEST_DEPENDENCIES)
	npm run test

DEV_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)
include $(PATH_TO_FUKOWL_MAKEFILES)/uisvr/dev.mk

include $(PATH_TO_FUKOWL_MAKEFILES)/uisvr/grpc.molecular.mk
PATH_TO_APISVR_GEN=$(PATH_TO_APISVR)/services/gen
PATH_TO_GRPC_SOURCE_DIR=$(PATH_TO_APISVR_GEN)/grpc
PATH_TO_GRPC_DEST_DIR=src/lib/server/protos
GRPC_RESOURCES=users channels chat_messages

include $(PATH_TO_FUKOWL_MAKEFILES)/uisvr/openapi.molecular.mk
PATH_TO_OPENAPI_TS_FILE=src/lib/openapi.d.ts

include $(PATH_TO_FUKOWL_MAKEFILES)/uisvr/test_integration.mk
