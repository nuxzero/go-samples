package main

import (
	"context"
	"log"
	"net"

	// import proto package
	pb "github.com/nuxzero/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// import grpc package
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v\n", err)
	}

	// Look up for the "blogdb" database if not exist, it will create one.
	// And look up for the "blog" collection if not exist, it will create one.
	collection = client.Database("blogdb").Collection("blog")

	if err != nil {
		log.Fatal(err)
	}

	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v/n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// Create a gRPC server
	s := grpc.NewServer()

	// Register the GreetService with the server
	pb.RegisterBlogServiceServer(s, &Server{})

	// Register the GreetService with the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
