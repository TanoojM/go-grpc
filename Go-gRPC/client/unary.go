package main

import (
	"context"
	"log"
	"time"

	pb "github.com/TanoojM/go-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// As you can see, we call the method on the stub we got earlier.
	// In our method parameters we create and populate a request protocol buffer object (in our case NoParam).
	// We also pass a context.Context object which lets us change our RPC’s behavior if
	// necessary, such as time-out/cancel an RPC in flight. If the call doesn’t return
	// an error, then we can read the response information from the server from the
	// first return value.
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
