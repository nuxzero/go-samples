package main

import (
	"context"
	"log"

	pb "github.com/nuxzero/grpc-go-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")

	req := &pb.BlogId{Id: id}
	_, err := c.DeleteBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was deleted!")
}
