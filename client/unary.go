package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Harsh-Bhabar/grpc-demo-proj/proto"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParams{})

	if err != nil {
		log.Fatalf("Cauld not read : %v\n", err)
	}

	log.Println("Response - ", res.Message)
}
