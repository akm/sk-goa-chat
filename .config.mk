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


# ## APP_PORT
#
# APP_ENV                        | APP_STAGE_TYPE           | apisvr HTTP | apisvr gRPC | ui HTTP  | mysql    | firebase authentication | swagger ui
# -------------------------------|--------------------------|-------------|-------------|----------|----------|-------------------------|------------
# server                         | staging,production       | 8000        | 8080        | 4173     | 3306     | ?                       |
# dev                            | local                    | 8000        | 8080        | 5173     | 3306     | 9099                    | 8090
# clients/uisvr/test:integration | staging,production       | 8000        | 8080        | 4173     | 3306     | ?                       |
# clients/uisvr/test:integration | local,github             | 8001        | 8081        | 4173     | 3307     | 9090                    |
# clients/uisvr/test:unit        | local,github             | -           | -           | -        | -        | -                       |
# servers/apisvr/test            | local,github             | -           | -           | -        | 3311     | 9091                    |

APP_PORT_APISVR_HTTP_dev?=8000
APP_PORT_APISVR_HTTP_e2e_test?=8001
APP_PORT_APISVR_GRPC_dev?=8080
APP_PORT_APISVR_GRPC_e2e_test?=8081
APP_PORT_UISVR_HTTP_dev?=5173
APP_PORT_UISVR_HTTP_e2e_test?=4173
APP_PORT_MYSQL_dev?=3306
APP_PORT_MYSQL_e2e_test?=3307
APP_PORT_MYSQL_unit_test?=3311
APP_PORT_FIREBASE_AUTH_dev?=9099
APP_PORT_FIREBASE_AUTH_e2e_test?=9090
APP_PORT_FIREBASE_AUTH_unit_test?=9091
APP_PORT_FIREBASE_EMULATOR_SUITE_dev?=4000
APP_PORT_SWAGGERUI_dev?=8090

## MySQL の設定

APP_MYSQL_DATABASE_NAME?=sk_goa_chat_db

## gcloud の設定

GOOGLE_CLOUD_PROJECT=sk-goa-chat

## Firebase の設定

APP_FIREBASE_API_KEY?=firebase-api-key-dummy1
