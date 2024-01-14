NODE_MODULES=node_modules
$(NODE_MODULES):
	$(MAKE) install

.PHONY: install
install:
	npm install

.PHONY: build
build: setup
	npm run build

.PHONY: lint
lint:
	npm run lint

.PHONY: dev
dev: setup
	npm run dev

.PHONY: preview
preview: setup
	npm run preview

.PHONY: test
test:
	npm run test

.PHONY: setup
setup: $(NODE_MODULES) 
	$(MAKE) $(FIREBASE_CONFIG_FILE)

FIREBASE_CONFIG_FILE=src/lib/firebase/firebaseconfig.ts
$(FIREBASE_CONFIG_FILE):
	cp $(FIREBASE_CONFIG_FILE).dummy $(FIREBASE_CONFIG_FILE)
