PATH_TO_UISVR:=../..
PATH_TO_PROJECT:=$(PATH_TO_UISVR)/../..
PATH_TO_FUKOWL_ROOT=$(PATH_TO_PROJECT)/vendor/fukowl
PATH_TO_FUKOWL_MAKEFILES=$(PATH_TO_FUKOWL_ROOT)/makefiles

include $(PATH_TO_PROJECT)/.config.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/locale2e/envs.mk
include $(PATH_TO_FUKOWL_MAKEFILES)/locale2e/servers.mk

# tmp ディレクトリはテスト中で使用する一時ファイルのために使用します
TMP_DIR=tmp
$(TMP_DIR):
	mkdir -p $(TMP_DIR)

.PHONY:
setup_tmp_dir:
	mkdir -p $(TMP_DIR)

.PHONY: test_setup
test_setup: $(TMP_DIR)
