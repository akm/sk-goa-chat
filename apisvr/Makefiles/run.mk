HTTP_PORT?=8000
GRPC_PORT?=8080

APISVR_OPTIONS=\
	-http-port $(HTTP_PORT) \
	-grpc-port $(GRPC_PORT)

.PHONY: run
run:
	go run ./services/cmd/apisvr $(APISVR_OPTIONS)

.PHONY: run_with_debug
run_with_debug:
	go run ./services/cmd/apisvr $(APISVR_OPTIONS) -debug
