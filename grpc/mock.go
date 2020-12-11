package grpc

import (
	"github.com/spyzhov/grpc_mock/mock"
	"github.com/spyzhov/grpc_mock/protob"
	"google.golang.org/grpc"
)

// TODO: autogenerate
func Mock(srv *grpc.Server) {
	protob.RegisterGreeterServer(srv, new(mock.GreeterServer))
}
