.PHONY: servers_apisvr_services_gen
servers_apisvr_services_gen:
	FILE=servers/apisvr/services/gen/log/logger.go $(MAKE) single

.PHONY: servers_apisvr_services_cmd
servers_apisvr_services_cmd:
	FILE=servers/apisvr/services/cmd/apisvr/main.go $(MAKE) single && \
	FILE=servers/apisvr/services/cmd/apisvr/http.go $(MAKE) single

.PHONY: servers_biz_models
servers_biz_models:
	PATTERN=servers/biz/models/models.rb $(MAKE) match

.PHONY: single
single:
	ruby ./single.rb ../${FILE}

.PHONY: match
match:
	ruby ./match.rb ./${PATTERN}
