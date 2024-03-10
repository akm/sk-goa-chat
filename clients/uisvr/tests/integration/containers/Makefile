PATH_TO_PROJECT=../../../../..

include $(PATH_TO_PROJECT)/.config.mk
include ../envs.mk

CONTAINERS_ENVS=\
	PATH_TO_PROJECT=$(PATH_TO_PROJECT)

.PHONY: run
run:
	$(ENVS) $(CONTAINERS_ENVS) docker compose up

.PHONY: up
up:
	$(ENVS) $(CONTAINERS_ENVS) docker compose up -d

.PHONY: down
down:
	$(ENVS) $(CONTAINERS_ENVS) docker compose down
