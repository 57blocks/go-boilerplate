//+build wireinject

package httpsrv

import (
	"context"

	"github.com/google/wire"
	"gocloud.dev/aws/awscloud"
	"gocloud.dev/server"
)

func SetupAWS(ctx context.Context) (*server.Server, func(), error) {
	wire.Build(
		awscloud.AWS,
		applicationSet,
	)
	return nil, nil, nil
}
