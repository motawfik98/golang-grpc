package main

import (
	"context"
	pb "go-grpc-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
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

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Mohamed"},
		{FirstName: "Ahmed"},
		{FirstName: "Tawfik"},
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v\n", err)
	}
	log.Printf("LongGreet %s\n", res.Result)
}

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while calling GreetEveryone: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Mohamed"},
		{FirstName: "Ahmed"},
		{FirstName: "Tawfik"},
	}
	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending request: %v\n", req)
			stream.Send(req)
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
				log.Printf("error while receiving: %v\n", err)
				break
			}
			log.Printf("received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Mohamed"}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok { // grpc error
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			}
			log.Fatalf("unexpected gRPC error: %v\n", err)
		} else {
			log.Fatalf("Non gRPC error: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
