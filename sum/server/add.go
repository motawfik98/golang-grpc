package main

import (
	"context"
	pb "go-grpc-course/sum/proto"
	"log"
)

func (s *Server) Add(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Add method invoked with params: %v\n", in)

	return &pb.SumResponse{
		Sum: in.Num1 + in.Num2,
	}, nil
}
