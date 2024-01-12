package main

import (
	"context"
	"errors"
	"fmt"
	pb "go-grpc-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}
		log.Printf("receiving: %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
	return stream.SendAndClose(&pb.GreetResponse{Result: res})
}

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		res := fmt.Sprintf("Hello %s!", req.FirstName)
		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("error while sending data to client: %v\n", err)
		}
	}
	return nil
}

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with: %v\n", in)

	for i := 0; i < 3; i++ {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println("The client cancelled the request")
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{Result: "Hello" + in.FirstName}, nil
}
