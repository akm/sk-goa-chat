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
ENV_TYPE?=unit_test
else
ENV_TYPE?=dev
endif


MYSQL_DATABASE_NAME?=sk_goa_chat_db
