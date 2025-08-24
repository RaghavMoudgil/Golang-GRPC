package main

import (
	pb "Raghav_GRPC/proto"
	"fmt"
	"io"
	"log"
)

func (s *helloServer) SayHelloBiDirectionalStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println("this is the request ", req.Name)
		res := &pb.HelloResponse{
			Message: "hello bi-directional name" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending stream", err)
		}

	}
}
