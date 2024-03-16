PATH_TO_PROJECT:=../..
PATH_TO_FUKOWL_MAKEFILES=$(PATH_TO_PROJECT)/vendor/fukowl/makefiles

.PHONY: default
default: up

include $(PATH_TO_PROJECT)/.config.mk

SWAGGERUI_MAPPING_HOST_PATH?=$(PATH_TO_APISVR)/services/gen/http/openapi3.json
SWAGGERUI_MAPPING_CONTAINER_PATH?=/srv/apisvr/openapi3.json

ENVS=\
	SWAGGERUI_PORT=$(APP_PORT_SWAGGERUI_dev) \
	SWAGGERUI_MAPPING_HOST_PATH=$(SWAGGERUI_MAPPING_HOST_PATH) \
	SWAGGERUI_MAPPING_CONTAINER_PATH=$(SWAGGERUI_MAPPING_CONTAINER_PATH) \
	SWAGGERUI_JSON_CONTAINER_PATH=$(SWAGGERUI_MAPPING_CONTAINER_PATH)

.PHONY: envs
envs:
	@echo $(ENVS)

.PHONY: run
run:
	$(ENVS) docker compose up

.PHONY: up
up:
	$(ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS) docker compose down
