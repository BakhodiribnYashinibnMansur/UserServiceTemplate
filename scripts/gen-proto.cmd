protoc --proto_path=proto/crud_service --gofast_out=plugins=grpc:. crud.proto
protoc --proto_path=proto/data_service --proto_path=proto/ --gofast_out=plugins=grpc:. user.proto