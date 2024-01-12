package main

import (
	pb "go-grpc-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// to use reflection, open the terminal and type
	// evans --host localhost --port 50053 --reflection retl
	// make sure to install evans first
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}
