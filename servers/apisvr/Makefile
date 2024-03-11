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
include ./Makefiles/goa.mk
include ./Makefiles/test.mk

# run, run_with_devbug
include $(PATH_TO_FUKOWL_MAKEFILES)/apisvr/run.organism.mk
# go run に渡されるので先頭のドットが必要です
PATH_TO_CMD_APISVR=./services/cmd/apisvr

# dev, dev_container_up, dev_container_down
include $(PATH_TO_FUKOWL_MAKEFILES)/apisvr/dev.organism.mk
DEV_MYSQL_DSN?='$(shell $(MAKE) -C $(PATH_TO_MYSQL) --no-print-directory dsn)'

DEV_SERVER_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)" \
	APP_STAGE=$(APP_STAGE) \
	APP_MYSQL_DSN=$(DEV_MYSQL_DSN) \
	APP_LOG_CONSOLE_WRITER=true \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY)
