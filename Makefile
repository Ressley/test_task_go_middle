PHONY: generate
generate:
	mkdir -p pkg/eventBus_v1
	protoc --go_out=pkg/eventBus_v1 --go_opt=paths=source_relative \
		--go-grpc_out=pkg/eventBus_v1 --go-grpc_opt=paths=import \
		proto/service.proto
	mv pkg/eventBus_v1/github.com/ressley/test_task_go_middle/pkg/eventBus_v1/* pkg/eventBus_v1/
	mv pkg/eventBus_v1/proto/* pkg/eventBus_v1/
	rm -rf pkg/eventBus_v1/proto
	rm -rf pkg/eventBus_v1/github.com

PHONY: start-client
start-client:
	@go run -v ./http-grpc-client 

PHONY: start-server
start-server:
	@go run -v ./grpc-server 