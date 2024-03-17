# ## APP_STAGE_TYPE
#
# - production
# - staging
# - local
# - github
#
# APP_STAGE_TYPE のインスタンスとして STAGE が複数個存在し、リポジトリにそれらの違いをコミットしなければならない場合もありえます。
# 例えば staging に対して staging1, staging2 それぞれで設定が異なる場合が考えられます。
# 逆に local のインスタンスとして各開発者の環境はそれぞれ別物ですが、それらの違いをコミットする必要がなければ local1, local2
# というような Stage を登録する必要はありませんし、通常はそのように行うべきではありません。
# ただし APP_STAGE_TYPE として local_windows と  local_mac のような区別を行った方が良い場合もあります。
#
# ## APP_STAGE の例
#
# - production
# - staging1
# - staging2
# - local
# - github

ifeq ($(GITHUB_ACTIONS),true)
APP_STAGE_TYPE?=github
else
APP_STAGE_TYPE?=local
endif

# local, github では STAGE は APP_STAGE_TYPE と同一です。デプロイ対象である production や
# staging では STAGE は APP_STAGE_TYPE と異なり、staging1 などの具体的なステージ名を指定します。
ifeq ($(APP_STAGE_TYPE),local)
APP_STAGE?=$(APP_STAGE_TYPE)
else ifeq ($(APP_STAGE_TYPE),github)
APP_STAGE?=$(APP_STAGE_TYPE)
else
APP_STAGE?=$(APP_STAGE_TYPE)1
endif

## APP_ENV
#
# (正確に言えば process cluster type みたいな名前のはずですが分かりやすい名前じゃないので APP_ENV にしました)
#
# APP_ENV | 起動するプロセス群
# ---------|------------------
# - server | uisvr, apisvr
# - dev     | uisvr, apisvr, mysql, firebase_auth
# - e2e_test | uisvr, apisvr, mysql, firebase_auth, clients/uisvr/test/integration
# - unit_test | uisvr のテスト( clients/uisvr/test/integration 以外), servers 以下のテスト
#
# この環境変数は、Makefie で処理を開始する際に設定します。
# つまりmakeのターゲットの実行時にのみ決定するので、静的に指定するものではありません。

include $(PATH_TO_GOAVELTE_MAKEFILES)/default/ports.mk

## MySQL の設定

APP_MYSQL_DATABASE_NAME?=sk_goa_chat_db

## gcloud の設定

GOOGLE_CLOUD_PROJECT=sk-goa-chat

## Firebase の設定

APP_FIREBASE_API_KEY?=firebase-api-key-dummy1

## 各ディレクトリへのパス

# PATH_TO_PROJECT が各Makefile で設定されている前提
PATH_TO_SERVERS=$(PATH_TO_PROJECT)/servers
PATH_TO_APISVR=$(PATH_TO_SERVERS)/apisvr
PATH_TO_APPLIB=$(PATH_TO_SERVERS)/applib
PATH_TO_BIZ=$(PATH_TO_SERVERS)/biz
PATH_TO_SERVER_CONTAINERS=$(PATH_TO_SERVERS)/containers
PATH_TO_FB_EMU=$(PATH_TO_SERVER_CONTAINERS)/firebase-emulators
PATH_TO_MYSQL=$(PATH_TO_SERVER_CONTAINERS)/mysql
PATH_TO_LOCALDEV=$(PATH_TO_SERVER_CONTAINERS)/localdev
PATH_TO_LOCALTEST=$(PATH_TO_SERVER_CONTAINERS)/localtest
PATH_TO_DBMIGRATIONS=$(PATH_TO_SERVERS)/dbmigrations
PATH_TO_SERVER_REPLACEMENTS=$(PATH_TO_SERVERS)/replacements

PATH_TO_CLIENTS=$(PATH_TO_PROJECT)/clients
PATH_TO_UISVR=$(PATH_TO_CLIENTS)/uisvr
PATH_TO_TEST_INTEGRATION=$(PATH_TO_UISVR)/tests/integration
PATH_TO_TEST_INTEGRATION_CONTAINERS=$(PATH_TO_TEST_INTEGRATION)/containers

PATH_TO_SWAGGERUI=$(PATH_TO_PROJECT)/tools/swaggerui
