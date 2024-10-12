build_proto:
	protoc -I=. \
		--go_out=backend/proto --go_opt=paths=source_relative \
		--go-grpc_out=backend/proto --go-grpc_opt=paths=source_relative \
		docs.proto

	protoc -I=. \
		--js_out=import_style=commonjs,binary:frontend/src/proto \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:frontend/src/proto \
		docs.proto

	# Move any stray generated files to appropriate directories
	find . -maxdepth 1 -name "docs*.go" -exec mv {} backend/proto/ \;
	find . -maxdepth 1 -name "docs*.js" -exec mv {} frontend/proto/ \;
