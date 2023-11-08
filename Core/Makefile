PROTOBUF_DIR=proto
SOURCE_PROTO_DIR=src

v1_domain:
	python -m grpc_tools.protoc \
	-I$(PROTOBUF_DIR) \
 	--python_out=$(SOURCE_PROTO_DIR)/proto \
 	--grpc_python_out=$(SOURCE_PROTO_DIR)/proto \
	--pyi_out=$(SOURCE_PROTO_DIR)/proto \
	$(PROTOBUF_DIR)/v1/domain/*.proto \
	&& rm $(SOURCE_PROTO_DIR)/proto/v1/domain/model_pb2_grpc.py


v1_subdomain:
	python -m grpc_tools.protoc \
	-I$(PROTOBUF_DIR) \
 	--python_out=$(SOURCE_PROTO_DIR)/proto \
 	--grpc_python_out=$(SOURCE_PROTO_DIR)/proto \
	--pyi_out=$(SOURCE_PROTO_DIR)/proto \
	$(PROTOBUF_DIR)/v1/subdomain/*.proto \
	&& rm $(SOURCE_PROTO_DIR)/proto/v1/subdomain/model_pb2_grpc.py

v1_route:
	python -m grpc_tools.protoc \
	-I$(PROTOBUF_DIR) \
 	--python_out=$(SOURCE_PROTO_DIR)/proto \
 	--grpc_python_out=$(SOURCE_PROTO_DIR)/proto \
	--pyi_out=$(SOURCE_PROTO_DIR)/proto \
	$(PROTOBUF_DIR)/v1/route/*.proto \
	&& rm $(SOURCE_PROTO_DIR)/proto/v1/route/model_pb2_grpc.py


v1_project:
	python -m grpc_tools.protoc \
	-I$(PROTOBUF_DIR) \
 	--python_out=$(SOURCE_PROTO_DIR)/proto \
 	--grpc_python_out=$(SOURCE_PROTO_DIR)/proto \
	--pyi_out=$(SOURCE_PROTO_DIR)/proto \
	$(PROTOBUF_DIR)/v1/project/*.proto \
	&& rm $(SOURCE_PROTO_DIR)/proto/v1/project/model_pb2_grpc.py