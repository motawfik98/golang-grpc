package main

import (
	pb "go-grpc-course/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.CalculatorServiceServer
}

var addr = "0.0.0.0:50053"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}
