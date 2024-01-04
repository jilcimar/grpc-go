package server

import "hello/grpc/pb"

type Server struct {
	pb.UnimplementedHelloServer
}
