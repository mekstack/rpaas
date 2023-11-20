PROTOBUF_DIR=proto
SOURCE_PROTO_DIR=src

proto_objects:
	python -m grpc_tools.protoc \
	-I. \
 	--python_out=$(SOURCE_PROTO_DIR) \
 	--grpc_python_out=$(SOURCE_PROTO_DIR) \
	--pyi_out=$(SOURCE_PROTO_DIR) \
	$(PROTOBUF_DIR)/*.proto