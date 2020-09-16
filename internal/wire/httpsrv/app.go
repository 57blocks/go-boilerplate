package httpsrv

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.opencensus.io/trace"
	"gocloud.dev/server/health"

	"github.com/57blocks/go-boilerplate/internal/handler/ping"
)

var applicationSet = wire.NewSet(
	trace.AlwaysSample,
	appHealthChecks,
	resourceSet,
)

func appHealthChecks() ([]health.Checker, func()) {
	return []health.Checker{}, func() {}
}

type resources struct {
	ping *ping.Resource
}

var resourceSet = wire.NewSet(
	setupHandler,
	wire.Struct(new(resources), "*"),
	ping.New,
)

func setupHandler(rs *resources) http.Handler {
	r := gin.New()

	r.GET("/v1/ping", rs.ping.Ping)

	return r
}
