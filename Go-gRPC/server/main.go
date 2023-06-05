package main

import (
	"log"
	"net"

	pb "github.com/TanoojM/go-grpc/proto"
	"google.golang.org/grpc"
)

// To build and start a server, we:

// Specify the port we want to use to listen for client requests using:
// lis, err := net.Listen(...).
// Create an instance of the gRPC server using grpc.NewServer(...).
// Register our service implementation with the gRPC server.
// Call Serve() on the server with our port details to do a blocking wait until the process is killed or Stop() is called.

// pb - all the files related to you proto, including the files that were created for us

const (
	port = "8080"
)

// helloServer helps to register that service
type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	// creating a listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Faile to start the server %v", err)
	}
	// once port is set
	// create a grpc server in context with the GreetService in the proto file
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{}) // this is the server that is capable of running our GreetService
	log.Printf("Server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err)
	}

}
