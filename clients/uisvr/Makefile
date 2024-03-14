PATH_TO_PROJECT:=../..
PATH_TO_FUKOWL_ROOT=$(PATH_TO_PROJECT)/vendor/fukowl
PATH_TO_FUKOWL_MAKEFILES=$(PATH_TO_FUKOWL_ROOT)/makefiles

include $(PATH_TO_PROJECT)/.config.mk

.PHONY: default
default: build lint test

NODE_MODULES=node_modules
$(NODE_MODULES):
	$(MAKE) install

.PHONY: install
install:
	npm install

# TODO npm run build ではなく npm run check を使うように変更
.PHONY: build
build: setup
	npm run build

.PHONY: lint
lint:
	npm run lint

.PHONY: preview
preview: setup
	npm run preview

.PHONY: test
test:
	npm run test

.PHONY: setup
setup: $(NODE_MODULES)
	$(MAKE) $(PATH_TO_FIREBASE_CONFIG_FILE)

PATH_TO_FIREBASE_CONFIG_FILE=src/lib/firebase/firebaseconfig.ts
$(PATH_TO_FIREBASE_CONFIG_FILE):
	cp $(PATH_TO_FIREBASE_CONFIG_FILE).dummy $(PATH_TO_FIREBASE_CONFIG_FILE)

PATH_TO_APISVR_GEN=$(PATH_TO_APISVR)/services/gen

include ./Makefiles/dev.mk
include ./Makefiles/test_integration.mk

include $(PATH_TO_FUKOWL_MAKEFILES)/uisvr/grpc.molecular.mk
PATH_TO_GRPC_SOURCE_DIR=$(PATH_TO_APISVR_GEN)/grpc
PATH_TO_GRPC_DEST_DIR=src/lib/server/protos
GRPC_RESOURCES=users channels chat_messages

include $(PATH_TO_FUKOWL_MAKEFILES)/uisvr/openapi.molecular.mk
PATH_TO_OPENAPI_TS_FILE=src/lib/openapi.d.ts
