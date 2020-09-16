package grpcsrv

import (
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/57blocks/go-boilerplate/gen/go/service"
	"github.com/57blocks/go-boilerplate/internal/server/echo"
)

var applicationSet = wire.NewSet(
	newGrpcServer,
	serviceSet,
)

type services struct {
	echo service.EchoServiceServer
}

var serviceSet = wire.NewSet(
	wire.Struct(new(services), "*"),
	echo.New,
)

func newGrpcServer(srvs *services) *grpc.Server {
	srv := grpc.NewServer()

	service.RegisterEchoServiceServer(srv, srvs.echo)

	return srv
}
