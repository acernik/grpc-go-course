package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"log"

	pb "github.com/acernik/grpc-go-course/calculator/proto"
)

func doSqrt(c pb.CalculatorServiceClient, number int32) error {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: number})
	if err != nil {
		e, ok := status.FromError(err)
		if !ok {
			return err
		}

		return fmt.Errorf("gRPC error message: %s, code: %d", e.Message(), e.Code())
	}

	log.Printf("Result Sqrt: %.2f\n", res.Result)

	return nil
}
