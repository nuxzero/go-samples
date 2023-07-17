package main

import (
	"log"

	pb "github.com/nuxzero/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true // change to false to use insecure connection
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt" // Certificate Authority Trust Certificate
		// Create the client TLS credentials using the CA trust certificate
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}
		// Append the client credentials to the gRPC dial options
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// Create a connection to the server with TLS option
	conn, err := grpc.Dial(addr, opts...)

	// Create a connection to the server
	// By default, gPRC uses TLS to secure the connection.
	// But we have not configured TLS, so use the WithInsecure() option
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	// This will close the connection when the main() function ends
	defer conn.Close()

	// Create a greet service client
	c := pb.NewGreetServiceClient(conn)

	// Call the Greet function to send request to the server
	doGreet(c)

	// doGreetManyTimes(c)

	// doLongGreet(c)

	// doGreetEveryone(c)

	// doGreetWithDeadline(c, 1*time.Second)
}
