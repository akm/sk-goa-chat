HTTP_PORT?=8000
GRPC_PORT?=8080

.PHONY: run
run:
	go run ./services/cmd/apisvr \
		-http-port $(HTTP_PORT) \
		-grpc-port $(GRPC_PORT)
