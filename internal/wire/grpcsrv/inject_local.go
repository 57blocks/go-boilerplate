//+build wireinject

package grpcsrv

import (
	"context"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

func SetupLocal(ctx context.Context) (*grpc.Server, func(), error) {
	wire.Build(
		applicationSet,
	)
	return nil, nil, nil
}
