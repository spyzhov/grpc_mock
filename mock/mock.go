package mock

import (
	"github.com/spyzhov/grpc_mock/protob"
	"google.golang.org/grpc"
)

// TODO: autogenerate
func Mock(srv *grpc.Server) {
	protob.RegisterGreeterServer(srv, new(GreeterServer))
}
