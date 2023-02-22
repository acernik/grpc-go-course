package main

import (
	"context"
	pb "github.com/acernik/grpc-go-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) error {
	log.Println("doGreetManyTimes was invoked")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Nik",
	})
	if err != nil {
		return err
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		log.Printf("got message %s\n", msg.Result)
	}

	return nil
}
