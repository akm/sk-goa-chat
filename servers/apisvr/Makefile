.PHONY: default
default: build lint test

PATH_TO_PROJECT=../..
PATH_TO_FUKOWL_ROOT=$(PATH_TO_PROJECT)/vendor/fukowl
PATH_TO_FUKOWL_MAKEFILES=$(PATH_TO_FUKOWL_ROOT)/makefiles

GO_ROOT_PACKAGE=apisvr

include ../../.config.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/asdf/reshim.atom.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/git/check.atom.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/golang/build.atom.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/golang/lint.atom.mk
include ./Makefiles/dev.mk
include ./Makefiles/goa.mk
include ./Makefiles/run.mk
include ./Makefiles/test.mk
