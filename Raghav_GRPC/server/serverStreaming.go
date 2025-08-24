package main

import (
	pb "Raghav_GRPC/proto"
	"fmt"
	"strconv"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.Namelist, stream pb.GreetService_SayHelloServerStreamingServer) error {
	fmt.Println("this is the list of the names ", req.Names)
	for i, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "name " + strconv.Itoa(i) + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
