package main

import (
	"github.com/WithLin/skywalking-go/trace"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address         = "localhost:11800"
	applicationCode = "skywalking-go-default"
)

func main() {


	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	trace.RegisterApplication()
	trace.RegisterApplicationInstance()
	for ; ;  {
		trace.Heartbeat()
	}

	time.Sleep(3 * time.Second)
}
