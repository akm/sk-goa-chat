GOOGLE_CLOUD_PROJECT=sk-goa-chat
APP_MYSQL_DATABASE_NAME?=sk_goa_chat_db
APP_FIREBASE_API_KEY?=firebase-api-key-dummy1

PATH_TO_GOAVELTE_MAKEFILES=$(PATH_TO_PROJECT)/vendor/goavelte/makefiles
include $(PATH_TO_GOAVELTE_MAKEFILES)/default/app_stage.mk
include $(PATH_TO_GOAVELTE_MAKEFILES)/default/ports.mk
include $(PATH_TO_GOAVELTE_MAKEFILES)/default/directories.mk
