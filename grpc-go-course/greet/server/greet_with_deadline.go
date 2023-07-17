package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nuxzero/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Example deadline code
// If the server takes time more the deadline, the server will receive context.DeadlineExceeded error
func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked wth %v\n", in)

	// Delay the response for 3 seconds
	for i := 0; i < 3; i++ {
		// If the client set timeout less than 3 seconds it will receive context.DeadlineExceeded error
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
