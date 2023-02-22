package main

import (
	"context"
	pb "github.com/acernik/grpc-go-course/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) error {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.CalculatorRequest{
		FirstNum:  2,
		SecondNum: 3,
	},
	)
	if err != nil {
		return err
	}

	log.Printf("Sum: %d\n", res.Result)

	return nil
}
