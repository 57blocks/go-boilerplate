// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package gateway

import (
	"context"
	"go.opencensus.io/trace"
	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
	"google.golang.org/grpc"
)

// Injectors from inject_local.go:

func SetupLocal(ctx context.Context, conn *grpc.ClientConn) (*server.Server, func(), error) {
	handler, err := newHandler(ctx, conn)
	if err != nil {
		return nil, nil, err
	}
	logger := _wireLoggerValue
	v, cleanup := appHealthChecks()
	exporter := _wireExporterValue
	sampler := trace.AlwaysSample()
	defaultDriver := _wireDefaultDriverValue
	options := &server.Options{
		RequestLogger:         logger,
		HealthChecks:          v,
		TraceExporter:         exporter,
		DefaultSamplingPolicy: sampler,
		Driver:                defaultDriver,
	}
	serverServer := server.New(handler, options)
	return serverServer, func() {
		cleanup()
	}, nil
}

var (
	_wireLoggerValue        = requestlog.Logger(nil)
	_wireExporterValue      = trace.Exporter(nil)
	_wireDefaultDriverValue = &server.DefaultDriver{}
)
