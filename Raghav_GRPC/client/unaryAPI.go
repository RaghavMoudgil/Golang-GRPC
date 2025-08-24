package main

import (
	pb "Raghav_GRPC/proto"
	"context"
	"fmt"
	"log"
	"time"
)

func callSayhello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Sayhello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("error while calling form the client", err)
	}
	fmt.Println("this is the call ", res)
}
