.PHONY: apisvr_services_gen
apisvr_services_gen:
	FILE=apisvr/services/gen/log/logger.go $(MAKE) modify

.PHONY: apisvr_services_cmd
apisvr_services_cmd:
	FILE=apisvr/services/cmd/apisvr/main.go $(MAKE) modify && \
	FILE=apisvr/services/cmd/apisvr/http.go $(MAKE) modify

.PHONY: modify
modify: call_modifier go_fmt

.PHONY: call_modifier
call_modifier:
	ruby ./main.rb ../${FILE}

.PHONY: go_fmt
go_fmt:
	go fmt ../${FILE}
