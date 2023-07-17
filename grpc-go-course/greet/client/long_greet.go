package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nuxzero/grpc-go-course/greet/proto"
)

// Example client code for client streaming RPC
// It will send multiple request to server and receive a single response from server
// This is the function that will be called by the client
func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")

	// Create a slice of request
	reqs := []*pb.GreetRequest{
		{FirstName: "Nux"},
		{FirstName: "John"},
		{FirstName: "Doe"},
	}

	// Create a stream by invoking the client
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		// Send request to server
		stream.Send(req)
		// Delay 1 second
		time.Sleep(1 * time.Second)
	}

	// Close and receive response from server
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
