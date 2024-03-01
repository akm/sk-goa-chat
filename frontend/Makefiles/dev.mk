.PHONY: dev
dev: setup
	npm run dev

DEV_ENVS=\
	GOOGLE_CLOUD_PROJECT=$(GOOGLE_CLOUD_PROJECT) \
	APP_FIREBASE_API_KEY=$(APP_FIREBASE_API_KEY) \
	FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)" \
	VITE_FIREBASE_AUTH_EMULATOR_HOST="127.0.0.1:$(APP_PORT_FIREBASE_AUTH_dev)" \
	VITE_UISVR_ORIGIN=http://localhost:$(APP_PORT_UISVR_HTTP_dev)

.PHONY: dev_backend
dev_backend:
	$(DEV_ENVS) APP_SKIP_DB_SCHEMA_DUMP=true \
		$(MAKE) -C ../backend dev

.PHONY: dev_run
dev_run:
	$(DEV_ENVS) vite dev

.PHONY: dev_containers_up_with_migration
dev_containers_up_with_migration:
	$(MAKE) -C ../backend/apisvr dev_containers_up_with_migration

.PHONY: dev_containers_down
dev_containers_down:
	$(MAKE) -C ../backend/apisvr dev_containers_down
