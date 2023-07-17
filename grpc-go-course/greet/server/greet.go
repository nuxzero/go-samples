package main

import (
	"context"
	"log"

	pb "github.com/nuxzero/grpc-go-course/greet/proto"
)

// Implement Greet function from GreetServiceServer interface that defined in greet_grp.pb.go file.
// This is the function that will be called by the client
// This is unary RPC (one request, one response) for server
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
