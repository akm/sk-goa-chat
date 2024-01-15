# See docs/stages.md for details.

# local以外では環境変数 STAGE を設定し、変数 STAGE のデフォルト値を上書きします。
# STAGE=github
# STAGE=staging
# STAGE=production
# STAGE=local

# ENV_TYPE=server
# ENV_TYPE=e2e_test
# ENV_TYPE=unit_test
# ENV_TYPE?=dev

ifeq ($(GITHUB_ACTIONS),true)
STAGE?=github
ENV_TYPE?=unit_test
else
STAGE?=local
ENV_TYPE?=dev
endif

STAGE_ENV_production_server=prodsrv
STAGE_ENV_staging_server=stgsrv
STAGE_ENV_local_dev=localdev
STAGE_ENV_local_e2e_test=locale2e
STAGE_ENV_local_unit_test=localtest
STAGE_ENV_github_e2e_test=githube2e
STAGE_ENV_github_unit_test=githubtest
STAGE_ENV?=$(STAGE_ENV_$(STAGE)_$(ENV_TYPE))

.PHONY: stage_env
stage_env:
	@echo $(STAGE_ENV)

MYSQL_DATABASE_NAME?=sk_goa_chat_db
