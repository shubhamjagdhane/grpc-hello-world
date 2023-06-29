package main

import (
	"context"
	"log"

	pb "github.com/shubhamjagdhane/grpc-hello-world/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoke with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello, " + in.FirstName,
	}, nil
}
