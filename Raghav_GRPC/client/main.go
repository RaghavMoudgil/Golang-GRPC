package main

import (
	pb "Raghav_GRPC/proto"
	"fmt"
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
	// fmt.Println("------------------server stream -------------------")

	// callSayhelloServerStraming(client, name) //calling server streaming
	// fmt.Println("-------------end-------------------")
	// fmt.Println("-------------client stream-------------------")

	// callSayhelloClientStraming(client, name) //colling client straming
	fmt.Println("------------------end -------------------")
	fmt.Println("------------------Bi-directional stram-------------------")
	callSayhelloBiDirectionalStraming(client, name)
	fmt.Println("------------------end -------------------")

}
