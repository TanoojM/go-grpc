package main

import (
	"log"

	pb "github.com/TanoojM/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = "8080"
)

// first connecting to the same port as server for communication

func main() {

	// To call service methods, we first need to create a gRPC channel to communicate
	// with the server. We create this by passing the server address and port number
	// to grpc.Dial()
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	// Creating a stub
	// Once the gRPC channel is setup, we need a client stub to perform RPCs.
	// We get it using the NewRouteGuideClient method provided by the pb package
	// generated from the .proto file.
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Tany", "Srik", "sai", "Vishnu"},
	}

	// unary API
	callSayHello(client)

	// server streaming
	callSayHelloServerStream(client, names)

	// client streaming
	callSayHelloClientStream(client, names)

	// bi direct streaming
	callSayHelloBidirectionalStream(client, names)

}
