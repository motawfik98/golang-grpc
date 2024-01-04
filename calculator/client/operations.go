package main

import (
	"context"
	pb "go-grpc-course/calculator/proto"
	"io"
	"log"
)

func doAddition(client pb.CalculatorServiceClient) {
	res, err := client.Sum(context.Background(), &pb.SumRequest{
		Num1: 3,
		Num2: 10,
	})

	if err != nil {
		log.Fatalf("Failed to add: %v\n", err)
	}
	log.Printf("Total is: %d\n", res.Sum)
}

func doPrime(client pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{Number: 120}
	stream, err := client.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Prime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading from the stream: %v\n", err)
		}
		log.Printf("Prime: %d", msg.Factor)
	}

}
