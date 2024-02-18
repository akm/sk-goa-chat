.PHONY: backend_apisvr_services_gen
backend_apisvr_services_gen:
	FILE=backend/apisvr/services/gen/log/logger.go $(MAKE) single

.PHONY: backend_apisvr_services_cmd
backend_apisvr_services_cmd:
	FILE=backend/apisvr/services/cmd/apisvr/main.go $(MAKE) single && \
	FILE=backend/apisvr/services/cmd/apisvr/http.go $(MAKE) single

.PHONY: single
single:
	ruby ./single.rb ../${FILE}
