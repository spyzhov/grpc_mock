

generate:
	@protoc --proto_path=proto --go_out=plugins=grpc,paths=source_relative:protob proto/*.proto
	@protoc --proto_path=proto --descriptor_set_out=service.protoset --include_imports proto/*.proto
