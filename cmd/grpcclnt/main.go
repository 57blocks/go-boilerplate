package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/57blocks/go-boilerplate/gen/go/resource"
	"github.com/57blocks/go-boilerplate/gen/go/service"
)

var (
	address string
)

func init() {
	flag.StringVar(&address, "address", "localhost:50051", "address of the server")
}

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := service.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &resource.StringMessage{Value: "hello grpc"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Echo: %s", r.GetValue())
}
