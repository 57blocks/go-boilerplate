package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/57blocks/go-boilerplate/internal/wire/httpsrv"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "8080", "port the http server listens to.")
}

// main function to boot up everything
func main() {
	flag.Parse()

	ctx := context.Background()
	srv, cleanup, err := httpsrv.SetupLocal(ctx)
	if err != nil {
		log.Fatalf("failed to setup the server, got error: %s", err)
	}
	defer cleanup()

	// Listen and serve HTTP.
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	log.Printf("starting server listening at :%s ...", port)
	log.Fatal(srv.ListenAndServe(":" + port))
}
