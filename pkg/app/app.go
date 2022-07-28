package app

import (
	"context"
	v1 "rabbitmqdddv2/pkg/app/controllers/v1"
	"rabbitmqdddv2/pkg/domain"
	"rabbitmqdddv2/pkg/infra"
	"rabbitmqdddv2/pkg/infra/logger"
	"time"

	trace "github.com/BetaLixT/appInsightsTrace"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var dependencySet = wire.NewSet(
	logger.NewLoggerFactory,
	domain.DependencySet,
	infra.DependencySet,
	v1.NewRabbitMqController,
	NewApp,
)

func Start() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}
	app.startService("<ROUTING_KEY>")
}

type app struct {
	lgr      *logger.LoggerFactory
	v1fcast  *v1.RabbitMqController
	insights *trace.AppInsightsCore
}

func NewApp(
	lgr *logger.LoggerFactory,
	v1fcast *v1.RabbitMqController,
	insights *trace.AppInsightsCore,
) *app {
	return &app{
		lgr:      lgr,
		v1fcast:  v1fcast,
		insights: insights,
	}
}

func (a *app) startService(topic string) {

	// Calling Listener Function
	a.v1fcast.Csub(a.lgr.NewLogger(), topic)
}

func (a *app) traceRequest(
	context context.Context,
	method,
	path,
	query,
	agent,
	ip string,
	status,
	bytes int,
	start,
	end time.Time) {
	latency := end.Sub(start)

	lgr := a.lgr.NewLogger()
	a.insights.TraceRequest(
		context,
		method,
		path,
		query,
		status,
		bytes,
		ip,
		agent,
		start,
		end,
		map[string]string{},
	)
	lgr.Info(
		"Request",
		zap.Int("status", status),
		zap.String("method", method),
		zap.String("path", path),
		zap.String("query", query),
		zap.String("ip", ip),
		zap.String("userAgent", agent),
		zap.Time("mvts", end),
		zap.String("pmvts", end.Format("2006-01-02T15:04:05-0700")),
		zap.Duration("latency", latency),
		zap.String("pLatency", latency.String()),
	)
}
