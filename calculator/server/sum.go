package main

import (
	"context"
	"log"

	pb "github.com/acernik/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, cr *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Sum function invoked with %v\n", cr)

	return &pb.CalculatorResponse{
		Result: cr.FirstNum + cr.SecondNum,
	}, nil
}
