package main

import (
	"context"
	pb "go-grpc-course/calculator/proto"
	"io"
	"log"
	"time"
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

func doAvg(client pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := client.Avg(context.Background())
	if err != nil {
		log.Fatalf("error while calling Avg: %v\n", err)
	}

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	for _, req := range reqs {
		log.Printf("sending Avg request: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("cannot receive response from Avg: %v\n", err)
	}
	log.Printf("result: %f\n", res.Result)
}
