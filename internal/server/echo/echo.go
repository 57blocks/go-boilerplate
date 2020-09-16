package echo

import (
	"context"
	"fmt"

	"github.com/57blocks/go-boilerplate/gen/go/resource"
	"github.com/57blocks/go-boilerplate/gen/go/service"
)

// New creates an server.
func New() service.EchoServiceServer {
	return &defaultEchoService{origin: "default echo service"}
}

type defaultEchoService struct {
	service.UnimplementedEchoServiceServer
	origin string
}

func (s *defaultEchoService) Echo(_ context.Context, in *resource.StringMessage) (*resource.StringMessage, error) {
	return &resource.StringMessage{
		Value: fmt.Sprintf("%s, from: %s", in.GetValue(), s.origin),
	}, nil
}
