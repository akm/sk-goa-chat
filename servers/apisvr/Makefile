.PHONY: default
default: build lint test

PROJECT_ROOT=../..
FUKOWL_ROOT=$(PROJECT_ROOT)/vendor/fukowl
FUKOWL_MAKEFILES=$(FUKOWL_ROOT)/makefiles

GO_ROOT_PACKAGE=apisvr

include ../../.config.mk
include $(FUKOWL_MAKEFILES)/asdf/reshim.atom.mk
include $(FUKOWL_MAKEFILES)/git/check.atom.mk
include $(FUKOWL_MAKEFILES)/golang/build.atom.mk
include $(FUKOWL_MAKEFILES)/golang/lint.atom.mk
include ./Makefiles/dev.mk
include ./Makefiles/goa.mk
include ./Makefiles/run.mk
include ./Makefiles/sqlboiler.mk
include ./Makefiles/test.mk
