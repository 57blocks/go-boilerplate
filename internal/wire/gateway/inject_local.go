//+build wireinject

package gateway

import (
	"context"

	"github.com/google/wire"
	"go.opencensus.io/trace"
	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
	"google.golang.org/grpc"
)

func SetupLocal(ctx context.Context, conn *grpc.ClientConn) (*server.Server, func(), error) {
	wire.Build(
		wire.InterfaceValue(new(requestlog.Logger), requestlog.Logger(nil)),
		wire.InterfaceValue(new(trace.Exporter), trace.Exporter(nil)),
		server.Set,
		applicationSet,
	)
	return nil, nil, nil
}
