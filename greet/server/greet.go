package main

import (
	"context"
	"fmt"
	pb "go-grpc-course/greet/proto"
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
