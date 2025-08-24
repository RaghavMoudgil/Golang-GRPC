package main

import (
	pb "Raghav_GRPC/proto"
	"context"
)

func (s *helloServer) Sayhello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Uniary API executed",
	}, nil
}
