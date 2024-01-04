package main

import (
	"context"
	"google.golang.org/grpc"
	"hello/grpc/pb"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Ola, " + in.GetName()}, nil
}

func main() {
	println("Rodando o servidor gRPC")

	listener, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Falha no servidor: %v", err)
	}
}
