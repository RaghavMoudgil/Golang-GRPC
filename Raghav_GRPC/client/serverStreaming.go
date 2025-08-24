package main

import (
	pb "Raghav_GRPC/proto"
	"context"
	"fmt"
	"io"
	"log"
)

func callSayhelloServerStraming(client pb.GreetServiceClient, names *pb.Namelist) {

	fmt.Print("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("error while sending name at server side", err)

	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming", err)

		}
		fmt.Println(message)
	}
	fmt.Print("streaming finished")
}
