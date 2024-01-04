package main

import (
	"context"
	"fmt"
	pb "go-grpc-course/greet/proto"
	"io"
	"log"
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
