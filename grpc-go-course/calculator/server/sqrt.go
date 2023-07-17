package main

import (
	"context"
	"log"
	"math"

	pb "github.com/nuxzero/grpc-go-course/calculator/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Example code of returning a gRPC error
func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with %v\n", in)

	number := in.Number

	if number < 0 {
		// Use status.Errorf from 'status' package to generate a gRPC error
		// You can choose gRPC error codes in the 'codes' package
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Received a negative number: %d\n",
			number,
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
