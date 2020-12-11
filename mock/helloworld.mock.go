package mock

import (
	"context"

	"github.com/spyzhov/grpc_mock/manager"
	"github.com/spyzhov/grpc_mock/protob"
)

// TODO: autogenerate
type GreeterServer struct {
	protob.UnimplementedGreeterServer
}

// TODO: autogenerate
func (s *GreeterServer) SayHello(_ context.Context, in *protob.HelloRequest) (*protob.HelloReply, error) {
	ret, err := manager.Call("helloworld", "Greeter", "SayHello", in, new(protob.HelloReply))
	out, _ := ret.(*protob.HelloReply)
	return out, err
}
