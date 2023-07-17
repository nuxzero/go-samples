package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/nuxzero/grpc-go-course/greet/proto"
)

// Example of client code for bidirectional streaming RPC.
// This will send a stream of request to server and receive a stream of response from server.
// In order to wait for all the response from server, we need to create a goroutine.
// And remember to close the channel once all the response are received.
func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Nux"},
		{FirstName: "Zero"},
		{FirstName: "NuxZero"},
	}

	// Create a channel to wait for all responses from server
	waitc := make(chan struct{})

	// Create goroutine to send requests to server
	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			// Send request to server
			stream.Send(req)
			// Delay 1 second
			time.Sleep(1 * time.Second)
		}
		// Close communication once all requests are sent
		stream.CloseSend()
	}()

	// Create goroutine to receive responses from server
	go func() {
		for {
			// Receive response from server
			res, err := stream.Recv()

			// break if reached end of stream
			if err == io.EOF {
				break
			}

			// break if error
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		// Close channel once all responses are received
		close(waitc)
	}()

	// Wait until the client received all responses from server
	<-waitc
}
