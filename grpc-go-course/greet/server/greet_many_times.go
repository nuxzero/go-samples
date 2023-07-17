package main

import (
	"fmt"
	"log"

	pb "github.com/nuxzero/grpc-go-course/greet/proto"
)

// Send stream of data from server to client
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	// loop 10 times
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		// send response back to client
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	// no error
	return nil
}
