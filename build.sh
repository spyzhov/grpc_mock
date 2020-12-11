#!/bin/bash
set -e

if [[ -z "${PROTO_PATH}" ]]; then
  echo "env:PROTO_PATH was not set"
  exit 1
fi

if [[ -d "${PROTO_PATH}" ]]; then
  find "${PROTO_PATH}" -name \*.proto -exec cp {} proto/ \;
else
  echo "Directory ${PROTO_PATH} doesn't exists"
  exit 1
fi

if [[ "$(find proto/ -type f -name '*.proto' | wc -l)" == "0" ]]; then
  echo "At least one proto file should be set"
  exit 1
fi

echo "Protoc:Gen"
protoc --proto_path=proto --go_out=plugins=grpc,paths=source_relative:protob proto/*.proto
#protoc --proto_path=proto --descriptor_set_out=service.protoset --include_imports proto/*.proto
echo "Mock:Gen"
mock_gen -input="$(pwd)/protob" -output="$(pwd)/mock"

echo "Go:Modules"
go mod download
go mod vendor
echo "Go:Build"
go build -mod=vendor -o "$(go env GOPATH)/bin/grpc_mock" "."
echo "Go:Build:Done"
