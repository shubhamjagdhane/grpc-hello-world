package main

import (
	"context"
	"log"

	pb "github.com/shubhamjagdhane/grpc-hello-world/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoke\n")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Shubham"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Result)
}
