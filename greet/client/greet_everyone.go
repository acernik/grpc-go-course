package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/acernik/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) error {
	log.Println("doGreetEveryone function invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		return err
	}

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

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second * 1)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error receiving message: %v\n", err)
				break
			}

			log.Printf("Received message: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc

	return nil
}
