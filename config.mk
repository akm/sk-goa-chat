# STAGE
# - local
# - github
# - staging
# - production

# Environment type    | STAGE                    | apisvr HTTP | apisvr gRPC | ui HTTP  | mysql    | firebase authentication | メモ
# --------------------|--------------------------|-------------|-------------|----------|----------|-------------------------|--------------
# server              | staging,production       | 8000        | 8080        | 4173     | 3306     | 9099                    |
# e2e test            | staging,production       | 8000        | 8080        | 4173     | 3306     | 9099                    | 対象のサーバーにアクセスするので特にサーバーを起動したりはしない
# dev                 | local                    | 8000        | 8080        | 5173     | 3306     | 9099
# e2e test            | local,github             | 8001        | 8081        | 4173     | 3307     | 9090
# unit test / apisvr  | local,github             | -           | -           | -        | 3311     | 9091
# unit test / ui      | local,github             | -           | -           | -        | -        | -

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
