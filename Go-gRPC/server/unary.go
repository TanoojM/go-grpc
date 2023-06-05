package main

import (
	"context"

	pb "github.com/TanoojM/go-grpc/proto"
)

// The method is passed a context object for the RPC and the client’s NoParam
// protocol buffer request. It returns a HelloResponse protocol buffer object with the response
// information and an error. In the method we populate the HelloResponse with the appropriate
// information, and then return it along with a nil error to tell gRPC that we’ve
// finished dealing with the RPC and that the HelloResponse can be returned to the client.

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
