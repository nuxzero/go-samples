package main

import (
	"log"

	pb "github.com/nuxzero/grpc-go-course/blog/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	// Create a connection to the server
	// By default, gPRC uses TLS to secure the connection.
	// But we have not configured TLS, so use the WithInsecure() option
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	// This will close the connection when the main() function ends
	defer conn.Close()

	// Create a greet service client
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	readBlog(c, "NonExistingId")
	updateBlog(c, id)
	listBlogs(c)
	deleteBlog(c, id)
}
