package main

import (
	"log"
	"net"

	// import proto package
	pb "github.com/nuxzero/grpc-go-course/greet/proto"

	// import grpc package
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

// Server is used to implement the GreetServiceServer interface
type Server struct {
	pb.GreetServiceServer
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v/n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// Create a slice of server options
	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		// Create the TLS credentials from the cert file and the private key
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		// Append the credentials to the server options
		opts = append(opts, grpc.Creds(creds))
	}

	// Create a gRPC server
	// s := grpc.NewServer()
	// Create a gRPC server with TLS options
	s := grpc.NewServer(opts...)

	// Register the GreetService with the server
	pb.RegisterGreetServiceServer(s, &Server{})

	// Register the GreetService with the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
