package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"

	pb "github.com/acernik/grpc-go-course/calculator/proto"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with: %v\n", in)

	if in.Number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", in.Number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(in.Number)),
	}, nil
}
