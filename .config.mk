include $(PATH_TO_GOAVELTE_MAKEFILES)/default/app_stage.mk

include $(PATH_TO_GOAVELTE_MAKEFILES)/default/ports.mk

## MySQL の設定

APP_MYSQL_DATABASE_NAME?=sk_goa_chat_db

## gcloud の設定

GOOGLE_CLOUD_PROJECT=sk-goa-chat

## Firebase の設定

APP_FIREBASE_API_KEY?=firebase-api-key-dummy1

## 各ディレクトリへのパス

include $(PATH_TO_GOAVELTE_MAKEFILES)/default/directories.mk
