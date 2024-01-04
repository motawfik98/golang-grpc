package main

import (
	"context"
	pb "go-grpc-course/calculator/proto"
	"io"
	"log"
	"math"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Add method invoked with params: %v\n", in)

	return &pb.SumResponse{
		Sum: in.Num1 + in.Num2,
	}, nil
}

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Prime method was invoked with params: %v\n", in)

	var k int32 = 2
	num := in.Number
	for num > 1 {
		if num%k == 0 {
			stream.Send(&pb.PrimeResponse{Factor: k})
			num /= k
		} else {
			k++
		}
	}
	return nil
}

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was revoked")

	sum := int32(0)
	totalNums := int32(0)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving from Avg client: %v\n", err)
		}
		sum += msg.Number
		totalNums++
	}
	return stream.SendAndClose(&pb.AvgResponse{
		Result: float64(sum) / float64(totalNums),
	})
}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was revoked")

	currentMax := int32(-1)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("cannot receive message: %v\n", err)
		}
		currentMax = int32(math.Max(float64(currentMax), float64(req.Number)))
		res := &pb.MaxResponse{Result: currentMax}
		err = stream.Send(res)
		if err != nil {
			log.Fatalf("cannot send maximum: %v\n", err)
		}
	}
	return nil
}
