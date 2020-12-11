#!/bin/bash
set -e

if [[ "$(find "${PROTO_PATH}" -type f -name '*.proto' | wc -l)" == "0" ]]; then
  echo "At least one proto file should be set"
  exit 1
fi

echo "Protoc:Gen"
CMD="protoc"
for path in $PROTO_PATHS; do
  CMD="${CMD} --proto_path=${path}"
done
CMD="${CMD} --go_out=plugins=grpc,paths=source_relative:protob \"${PROTO_PATH}\"/*.proto"

eval "$CMD"
#protoc --proto_path=proto --descriptor_set_out=service.protoset --include_imports proto/*.proto
echo "Mock:Gen"
mock_gen -input="$(pwd)/protob" -output="$(pwd)/mock"

echo "Go:Modules"
go mod download
go mod vendor
echo "Go:Build"
go build -mod=vendor -o "$(go env GOPATH)/bin/grpc_mock" "."
echo "Go:Build:Done"
