.PHONY: default
default: build lint test

PATH_TO_PROJECT:=../..
include $(PATH_TO_PROJECT)/.config.mk

GO_ROOT_PACKAGE=biz

include $(PATH_TO_GOAVELTE_MAKEFILES)/asdf/reshim.mk
include $(PATH_TO_GOAVELTE_MAKEFILES)/git/check.mk
include $(PATH_TO_GOAVELTE_MAKEFILES)/golang/build.mk
include $(PATH_TO_GOAVELTE_MAKEFILES)/golang/lint.mk

TEST_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT)
include $(PATH_TO_GOAVELTE_MAKEFILES)/biz/test.mk

include $(PATH_TO_GOAVELTE_MAKEFILES)/biz/sqlboiler.mk
