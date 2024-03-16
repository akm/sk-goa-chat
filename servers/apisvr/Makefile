PATH_TO_PROJECT:=../..
PATH_TO_FUKOWL_ROOT=$(PATH_TO_PROJECT)/vendor/fukowl
PATH_TO_FUKOWL_MAKEFILES=$(PATH_TO_FUKOWL_ROOT)/makefiles

.PHONY: default
default: build lint test

GO_ROOT_PACKAGE=apisvr

include $(PATH_TO_PROJECT)/.config.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/asdf/reshim.atom.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/git/check.atom.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/golang/build.atom.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/golang/lint.atom.mk

# run, run_with_devbug
include $(PATH_TO_FUKOWL_MAKEFILES)/apisvr/run.mk
# go run に渡されるので先頭のドットが必要です
PATH_TO_CMD_APISVR=./services/cmd/apisvr

# dev, dev_container_up, dev_container_down
include $(PATH_TO_FUKOWL_MAKEFILES)/apisvr/dev.mk
DEV_SERVER_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)

include $(PATH_TO_FUKOWL_MAKEFILES)/apisvr/goa.mk
GOA_ROOT_PACKAGE=$(GO_ROOT_PACKAGE)
PATH_TO_GOA_SERVICES_DIR=./services

include $(PATH_TO_FUKOWL_MAKEFILES)/apisvr/test.mk
TEST_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)
