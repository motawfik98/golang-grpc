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

func doMax(client pb.CalculatorServiceClient) {
	log.Println("doMax function was revoked")

	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("cannot call Max: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		nums := []int32{1, 5, 3, 6, 2, 20}
		for _, num := range nums {
			log.Printf("sending num: %d\n", num)
			err := stream.Send(&pb.MaxRequest{Number: num})
			if err != nil {
				log.Printf("cannot send Max request: %v\n", err)
				break
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("failed to receive max: %v\n", err)
				break
			}
			log.Printf("current max is: %d\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
