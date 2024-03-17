include $(PATH_TO_GOAVELTE_MAKEFILES)/default/app_stage.mk

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

include $(PATH_TO_GOAVELTE_MAKEFILES)/default/directories.mk
