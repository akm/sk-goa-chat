ROOT_PATH=../../../../..

include $(ROOT_PATH)/.config.mk
include ../envs.mk

CONTAINERS_ENVS=\
	ROOT_PATH=$(ROOT_PATH)

.PHONY: run
run:
	$(ENVS) $(CONTAINERS_ENVS) docker compose up

.PHONY: up
up:
	$(ENVS) $(CONTAINERS_ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS) $(CONTAINERS_ENVS) docker compose down
