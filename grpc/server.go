package grpc

import (
	"log"
	"net"
	"strconv"

	"github.com/spyzhov/grpc_mock/mock"
	"google.golang.org/grpc"
)

func ListenAndServe(port int) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	mock.Mock(srv)
	log.Printf("gPRC: Serv on :%d", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
