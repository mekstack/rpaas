compile_proto:
	protoc \
	-I proto .\proto\objects.proto \
	--go_out=.\proto\pb \
	--go_opt=paths=source_relative \
	--go-grpc_out=.\proto\pb \
	--go-grpc_opt=paths=source_relative
	