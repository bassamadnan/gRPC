.PHONY: build_proto build run_servers clean

build_proto:
	protoc --go_out=. --go-grpc_out=. pkg/api/knn.proto

build:
	go build -o bin/server cmd/server/main.go

run_servers:
	@if [ -z "$(SERVERS)" ]; then \
		echo "Usage: make run_servers SERVERS=<number_of_servers>"; \
		exit 1; \
	fi
	@for i in $$(seq 0 $$(($(SERVERS)-1))); do \
		./bin/server -server=$$i -total=$(SERVERS) & \
		echo $$! >> .server_pids; \
	done
	@echo "Started $(SERVERS) servers. Use 'make clean' to stop all servers."

clean:
	@if [ -f .server_pids ]; then \
		echo "Stopping servers..."; \
		kill $$(cat .server_pids) 2>/dev/null || true; \
		rm .server_pids; \
		echo "All servers stopped."; \
	else \
		echo "No running servers found."; \
	fi
	@rm -f bin/server
