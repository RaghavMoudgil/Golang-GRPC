package main

import (
	"fmt"
	"net"

	pb "Raghav_GRPC/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.UnimplementedGreetServiceServer
}

func main() {
	lisn, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("error while starting server", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lisn); err != nil {
		fmt.Println("Error while starting GRPC server", err)
	}
}
