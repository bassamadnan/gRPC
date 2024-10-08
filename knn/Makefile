build_proto:
	protoc --go_out=. --go-grpc_out=. pkg/api/knn.proto

run_servers:
	@if [ -z "$(SERVERS)" ]; then \
		echo "Usage: make run_servers SERVERS=<number_of_servers>"; \
		exit 1; \
	fi
	@for i in $$(seq 0 $$(($(SERVERS)-1))); do \
		go run cmd/server/main.go -server $$i -total $(SERVERS) -k 3 & \
	done
	@echo "Started $(SERVERS) servers. Press Ctrl+C to stop all servers."
	@wait
