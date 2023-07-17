package main

import (
	"log"
	"net"

	// import proto package
	pb "github.com/nuxzero/grpc-go-course/calculator/proto"

	// import grpc package
	"google.golang.org/grpc"

	// import reflection package
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v/n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// Create a gRPC server
	s := grpc.NewServer()

	// Register the GreetService with the server
	pb.RegisterCalculatorServiceServer(s, &Server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	// Register the GreetService with the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
