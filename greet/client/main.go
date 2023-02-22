package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"

	pb "github.com/acernik/grpc-go-course/greet/proto"
)

const addr = "localhost:50051"

func main() {
	opts := make([]grpc.DialOption, 0)
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			panic(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//err = doGreet(c)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = doGreetManyTimes(c)
	//if err != nil {
	//	panic(err)
	//}

	//err = doLongGreet(c)
	//if err != nil {
	//	panic(err)
	//}

	//err = doGreetEveryone(c)
	//if err != nil {
	//	panic(err)
	//}

	err = doGreetWithDeadline(c, time.Second*5)
	if err != nil {
		panic(err)
	}
}
