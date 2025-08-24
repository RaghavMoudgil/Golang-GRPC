package main

import (
	pb "Raghav_GRPC/proto"
	"fmt"
	"io"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var message []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: message})
		}
		if err != nil {
			return err
		}
		fmt.Println("request recived with name", req.Name)
		message = append(message, "Hello", req.Name)
	}

}
