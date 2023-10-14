package main

import (
	"fmt"
	"log"

	pb "github.com/Harsh-Bhabar/grpc-demo-proj/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	fmt.Println("Starting client...")

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := pb.NamesList{
		Names: []string{"Harsh", "Mrinal", "Ayush"},
	}

	// Unary Streaming
	// CallSayHello(client)

	// server streaming
	CallSayHalloServerStreaming(client, names)

}
