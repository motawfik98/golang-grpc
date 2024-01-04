package main

import (
	"context"
	pb "go-grpc-course/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Add method invoked with params: %v\n", in)

	return &pb.SumResponse{
		Sum: in.Num1 + in.Num2,
	}, nil
}
