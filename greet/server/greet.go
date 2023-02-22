package main

import (
	"context"
	"log"

	pb "github.com/acernik/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, gr *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function invoked with %v\n", gr)

	return &pb.GreetResponse{
		Result: "Hello " + gr.FirstName,
	}, nil
}
