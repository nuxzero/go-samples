package main

import (
	"io"
	"log"

	pb "github.com/nuxzero/grpc-go-course/greet/proto"
)

// Example of server code for bidirectional streaming RPC.
// Server will read a stream of request from client.
// And send back a stream of response to client.
func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		// Receive request from client
		req, err := stream.Recv()

		// if reached end of stream
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello" + req.FirstName + "!"
		// Send response to client
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
