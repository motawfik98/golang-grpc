package main

import (
	"context"
	pb "go-grpc-course/greet/proto"
	"io"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Tawfik"})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{FirstName: "Tawfik"}
	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GreetManyTimes %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %s", msg.Result)
	}
}
