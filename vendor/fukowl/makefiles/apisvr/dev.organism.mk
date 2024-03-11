.PHONY: dev
dev: dev_containers_up
	$(DEV_SERVER_ENVS) $(MAKE) run_with_debug

.PHONY: dev_containers_up
dev_containers_up:
	$(DEV_SERVER_ENVS) $(MAKE) -C $(PATH_TO_LOCALDEV) up

.PHONY: dev_containers_down
dev_containers_down:
	$(DEV_SERVER_ENVS) $(MAKE) -C $(PATH_TO_LOCALDEV) down
