package main

import (
	"context"
	pb "go-grpc-course/calculator/proto"
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
