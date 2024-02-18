.PHONY: backend_apisvr_services_gen
backend_apisvr_services_gen:
	FILE=backend/apisvr/services/gen/log/logger.go $(MAKE) single

.PHONY: backend_apisvr_services_cmd
backend_apisvr_services_cmd:
	FILE=backend/apisvr/services/cmd/apisvr/main.go $(MAKE) single && \
	FILE=backend/apisvr/services/cmd/apisvr/http.go $(MAKE) single

.PHONY: backend_biz_models
backend_biz_models:
	PATTERN=backend/biz/models/models.rb $(MAKE) match

.PHONY: single
single:
	ruby ./single.rb ../${FILE}

.PHONY: match
match:
	ruby ./match.rb ./${PATTERN}
