package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"

	"github.com/57blocks/go-boilerplate/internal/wire/gateway"
)

var (
	address string
	port    string
)

func init() {
	flag.StringVar(&address, "address", "localhost:50051", "address of the server")
	flag.StringVar(&port, "port", "8081", "port the gRPC gateway listens to.")
}

func main() {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	srv, cleanup, err := gateway.SetupLocal(context.Background(), conn)
	if err != nil {
		log.Fatalf("failed to setup the gateways, got error: %s", err)
	}
	defer cleanup()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("starting gateway listening at :%s ...", port)
	log.Fatal(srv.ListenAndServe(":" + port))
}
