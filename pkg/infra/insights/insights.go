package insights

import (
	"rabbitmqdddv2/pkg/infra/logger"

	trace "github.com/BetaLixT/appInsightsTrace"
)

func NewInsights(
	optn *trace.AppInsightsOptions,
	lgrf *logger.LoggerFactory,
) *trace.AppInsightsCore {
	lgr := lgrf.NewLogger()
	return trace.NewAppInsightsCore(
		optn,
		&traceExtractor{},
		lgr,
	)
}
