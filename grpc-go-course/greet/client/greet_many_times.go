package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nuxzero/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Nux",
	}

	// call the RPC with the request
	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		// receive stream
		msg, err := stream.Recv()

		// if end of stream then stop looping
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Response from GreetManyTimes: %v\n", msg.Result)
	}
}
