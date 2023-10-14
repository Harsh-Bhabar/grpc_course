package main

import (
	"context"

	pb "github.com/Harsh-Bhabar/grpc-demo-proj/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
