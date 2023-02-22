package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"log"
	"time"

	pb "github.com/acernik/grpc-go-course/greet/proto"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) error {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Nik",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if !ok {
			return err
		}

		return fmt.Errorf("timeout exceeded with error message: %s and error code: %d", e.Message(), e.Code())
	}

	log.Printf("Response: %s\n", res.Result)

	return nil
}
