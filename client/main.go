package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hello/grpc/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{
		Name: "Jilcimar",
	}

	res, err := client.SayHello(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
