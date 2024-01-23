include ../../config.mk

SWAGGERUI_PORT?=8090
SWAGGERUI_JSON_PATH?=/srv/apisvr/services/gen/http/openapi3.json

ENVS=\
	SWAGGERUI_PORT=$(SWAGGERUI_PORT) \
	SWAGGERUI_JSON_PATH=$(SWAGGERUI_JSON_PATH)

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
