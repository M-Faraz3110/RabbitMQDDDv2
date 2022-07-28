package app

import (
	v1 "rabbitmqdddv2/pkg/app/controllers/v1"
	"rabbitmqdddv2/pkg/domain"
	"rabbitmqdddv2/pkg/infra"
	"rabbitmqdddv2/pkg/infra/logger"

	trace "github.com/BetaLixT/appInsightsTrace"
	"github.com/google/wire"
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

// func (a *app) traceRequest(
// 	context context.Context,
// 	method,
// 	path,
// 	query,
// 	agent,
// 	ip string,
// 	status,
// 	bytes int,
// 	start,
// 	end time.Time) {
// }
