//+build wireinject

package httpsrv

import (
	"context"

	"github.com/google/wire"
	"gocloud.dev/gcp/gcpcloud"
	"gocloud.dev/server"
)

func SetupGCP(ctx context.Context) (*server.Server, func(), error) {
	wire.Build(
		gcpcloud.GCP,
		applicationSet,
	)
	return nil, nil, nil
}
