package main

import (
	"os"
	"strconv"

	grpc "github.com/spyzhov/grpc_mock/grpc"
	"github.com/spyzhov/grpc_mock/manager"
)

//go:generate protoc --proto_path=proto --go_out=plugins=grpc,paths=source_relative:protob proto/helloworld.proto
//go:generate protoc --proto_path=proto --descriptor_set_out=service.protoset --include_imports proto/helloworld.proto
func main() {
	go manager.ListenAndServe(port("MANAGE_PORT", 8000))
	go grpc.ListenAndServe(port("GRPC_PORT", 8090))
	select {}
}

func port(name string, def int) int {
	value := os.Getenv(name)
	if d, err := strconv.ParseInt(value, 10, 64); err != nil {
		return def
	} else {
		return int(d)
	}
}
