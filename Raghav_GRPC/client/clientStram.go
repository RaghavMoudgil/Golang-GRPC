package main

import (
	pb "Raghav_GRPC/proto"
	"context"
	"fmt"
	"log"
	"time"
)

func callSayhelloClientStraming(client pb.GreetServiceClient, names *pb.Namelist) {
	fmt.Println("client straming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("unable to start client stram", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatal("unable to send stram", err)
		}
		fmt.Println("send the request with name ", req.Name)
		time.Sleep(time.Second)
	}
	req, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("error while closing stream", err)
	}
	log.Printf("client straming finished")

	fmt.Println(req.Messages)
}
