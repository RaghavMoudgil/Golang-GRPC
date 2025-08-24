package main

import (
	pb "Raghav_GRPC/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error while starting client side server: %v", err)

	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	callSayhello(client) //calling unary API

	name := &pb.Namelist{
		Names: []string{"Raghav", "Moudgil"},
	}
	callSayhelloServerStraming(client, name) //calling server streaming
}
