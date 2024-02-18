.PHONY: backend_apisvr_services_gen
backend_apisvr_services_gen:
	FILE=backend/apisvr/services/gen/log/logger.go $(MAKE) modify

.PHONY: backend_apisvr_services_cmd
backend_apisvr_services_cmd:
	FILE=backend/apisvr/services/cmd/apisvr/main.go $(MAKE) modify && \
	FILE=backend/apisvr/services/cmd/apisvr/http.go $(MAKE) modify

.PHONY: modify
modify:
	ruby ./main.rb ../${FILE}
	
