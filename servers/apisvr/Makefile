PATH_TO_PROJECT:=../..
PATH_TO_GOAVELTE_MAKEFILES=$(PATH_TO_PROJECT)/vendor/goavelte/makefiles

.PHONY: default
default: build lint test

GO_ROOT_PACKAGE=apisvr

include $(PATH_TO_PROJECT)/.config.mk
include $(PATH_TO_GOAVELTE_MAKEFILES)/asdf/reshim.mk
include $(PATH_TO_GOAVELTE_MAKEFILES)/git/check.mk
include $(PATH_TO_GOAVELTE_MAKEFILES)/golang/build.mk
include $(PATH_TO_GOAVELTE_MAKEFILES)/golang/lint.mk

# run, run_with_devbug
include $(PATH_TO_GOAVELTE_MAKEFILES)/apisvr/run.mk
# go run に渡されるので先頭のドットが必要です
PATH_TO_CMD_APISVR=./services/cmd/apisvr

# dev, dev_container_up, dev_container_down
DEV_SERVER_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)
include $(PATH_TO_GOAVELTE_MAKEFILES)/apisvr/dev.mk

GOA_ROOT_PACKAGE=$(GO_ROOT_PACKAGE)
PATH_TO_GOA_SERVICES_DIR=./services
include $(PATH_TO_GOAVELTE_MAKEFILES)/apisvr/goa.mk

TEST_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)
include $(PATH_TO_GOAVELTE_MAKEFILES)/apisvr/test.mk
