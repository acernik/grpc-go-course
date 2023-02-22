package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/acernik/grpc-go-course/calculator/proto"
)

const addr = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		panic(err)
	}
}
