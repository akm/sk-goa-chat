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
include ./Makefiles/dev.mk
include ./Makefiles/goa.mk
include ./Makefiles/test.mk

# run, run_with_devbug
include $(PATH_TO_FUKOWL_MAKEFILES)/apisvr/run.organism.mk
# go run に渡されるので先頭のドットが必要です
PATH_TO_CMD_APISVR=./services/cmd/apisvr
