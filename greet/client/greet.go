package main

import (
	"context"
	pb "github.com/acernik/grpc-go-course/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) error {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Nik",
	})
	if err != nil {
		return err
	}

	log.Printf("Greeting: %s\n", res.Result)

	return nil
}
