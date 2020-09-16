// Package main is the starting point of the notification applicatioin.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/57blocks/go-boilerplate/internal/wire/grpcsrv"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "50051", "port the gRPC server listens to.")
}

// main function to boot up everything
func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv, cleanup, err := grpcsrv.SetupLocal(context.Background())
	if err != nil {
		log.Fatalf("failed to setup the server, got error: %s", err)
	}
	defer cleanup()

	log.Printf("starting server listening at :%s ...", port)
	log.Fatal(srv.Serve(lis))
}
