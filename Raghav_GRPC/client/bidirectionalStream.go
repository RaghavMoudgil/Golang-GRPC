package main

import (
	pb "Raghav_GRPC/proto"
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

func callSayhelloBiDirectionalStraming(client pb.GreetServiceClient, names *pb.Namelist) {
	fmt.Println("Bi-directional stram started")
	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while initializing Bidirectional stream")
	}
	// to create a stream from both sides we are going to use channels
	ch := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while reciving stream", err)
			}
			fmt.Println(message)

		}
		close(ch)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending stream", err)
		}

		time.Sleep(time.Second)
	}
	stream.CloseSend()
	<-ch

	fmt.Println("bi-directinal stream ended ")
}
