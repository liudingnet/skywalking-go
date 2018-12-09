package main

import (
	pb "../../proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address         = "localhost:11800"
	applicationCode = "skywalking-go-default"
)

func main() {
	application := pb.Application{
		ApplicationCode: applicationCode,
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewApplicationRegisterServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r, err := c.ApplicationCodeRegister(ctx, &application)
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}
	log.Printf("application key[%v], value[%v]", r.Application.Key, r.Application.Value)

	time.Sleep(3 * time.Second)
}
