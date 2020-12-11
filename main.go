package main

import (
	"os"
	"strconv"

	"github.com/spyzhov/grpc_mock/grpc"
	"github.com/spyzhov/grpc_mock/manager"
)

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
