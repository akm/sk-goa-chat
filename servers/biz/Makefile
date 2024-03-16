.PHONY: default
default: build lint test

PATH_TO_PROJECT:=../..
PATH_TO_FUKOWL_ROOT=$(PATH_TO_PROJECT)/vendor/fukowl
PATH_TO_FUKOWL_MAKEFILES=$(PATH_TO_FUKOWL_ROOT)/makefiles

GO_ROOT_PACKAGE=biz

include $(PATH_TO_PROJECT)/.config.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/asdf/reshim.atom.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/git/check.atom.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/golang/build.atom.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/golang/lint.atom.mk

include $(PATH_TO_FUKOWL_MAKEFILES)/biz/test.mk
TEST_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT)

include $(PATH_TO_FUKOWL_MAKEFILES)/biz/sqlboiler.mk
