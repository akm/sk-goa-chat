.PHONY: default
default: build lint test

PATH_TO_PROJECT:=../..
PATH_TO_FUKOWL_MAKEFILES=$(PATH_TO_PROJECT)/vendor/fukowl/makefiles

GO_ROOT_PACKAGE=applib

include $(PATH_TO_PROJECT)/.config.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/asdf/reshim.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/golang/build.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/golang/lint.mk

TEST_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT)
include $(PATH_TO_FUKOWL_MAKEFILES)/applib/test.mk
