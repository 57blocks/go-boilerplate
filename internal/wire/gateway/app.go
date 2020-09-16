package gateway

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opencensus.io/trace"
	"gocloud.dev/server/health"
	"google.golang.org/grpc"

	"github.com/57blocks/go-boilerplate/gen/go/service"
)

var applicationSet = wire.NewSet(
	trace.AlwaysSample,
	appHealthChecks,
	gatewaySet,
)

func appHealthChecks() ([]health.Checker, func()) {
	return []health.Checker{}, func() {}
}

var gatewaySet = wire.NewSet(
	newHandler,
)

func newHandler(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux := runtime.NewServeMux()

	if err := service.RegisterEchoServiceHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return mux, nil
}
