.PHONY: dev
dev: setup
	npm run dev

DEV_ENVS=\
	GOOGLE_CLOUD_PROJECT=sk-goa-chat \
	FIREBASE_AUTH_EMULATOR_HOST='127.0.0.1:9099' \
	VITE_FIREBASE_AUTH_EMULATOR_HOST='127.0.0.1:9099' \
	VITE_UISVR_ORIGIN=http://localhost:5173

.PHONY: dev_apisvr
dev_apisvr:
	$(DEV_ENVS) $(MAKE) -C ../backend/apisvr dev

.PHONY: dev_run
dev_run:
	$(DEV_ENVS) vite dev

.PHONY: dev_containers_up_with_migration
dev_containers_up_with_migration:
	$(MAKE) -C ../backend/apisvr dev_containers_up_with_migration

.PHONY: dev_containers_down
dev_containers_down:
	$(MAKE) -C ../backend/apisvr dev_containers_down
