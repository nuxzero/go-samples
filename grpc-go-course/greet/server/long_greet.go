package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/nuxzero/grpc-go-course/greet/proto"
)

// Example server code for client streaming RPC.
// Server will receive a stream of request (multiple requests) from client.
// Once the server received all the request, it will send back a single response to client.
func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res := ""

	for {
		// receive request from client
		req, err := stream.Recv()

		// if reached end of stream
		if err == io.EOF {
			// finished reading client stream and return response
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
