build_proto:
	protoc --go_out=. --go-grpc_out=. pkg/proto/rideshare.proto
	protoc --go_out=. --go-grpc_out=. pkg/proto/state.proto

cert:
	cd cert && ../gen.sh && cd ../

.PHONY: build_proto cert
