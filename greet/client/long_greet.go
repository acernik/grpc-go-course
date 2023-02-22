package main

import (
	"context"
	pb "github.com/acernik/grpc-go-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) error {
	log.Println("doLongGreet function invoked ")

	reqs := []*pb.GreetRequest{
		{
			FirstName: "Nik",
		},
		{
			FirstName: "Kolian",
		},
		{
			FirstName: "Kolianych",
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		return err
	}

	for i, req := range reqs {
		log.Printf("Sending request: %v id: %d\n", req.FirstName, i)
		stream.Send(req)
		time.Sleep(time.Second * 1)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("LongGreet result: %s\n", res.Result)

	return nil
}
