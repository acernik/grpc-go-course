package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/acernik/grpc-go-course/calculator/proto"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//err = doSum(c)
	//if err != nil {
	//	panic(err)
	//}

	err = doSqrt(c, 5)
	if err != nil {
		panic(err)
	}
}
