build_proto:
	protoc -I=. \
		--go_out=backend/proto --go_opt=paths=source_relative \
		--go-grpc_out=backend/proto --go-grpc_opt=paths=source_relative \
		docs.proto

	# Move any stray generated files to appropriate directories
	find . -maxdepth 1 -name "docs*.go" -exec mv {} backend/proto/ \;
