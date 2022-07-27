package infra

import (
	"rabbitmqdddv2/pkg/domain/tablestorage"
	"rabbitmqdddv2/pkg/infra/config"
	"rabbitmqdddv2/pkg/infra/db"
	"rabbitmqdddv2/pkg/infra/insights"
	"rabbitmqdddv2/pkg/infra/logger"
	"rabbitmqdddv2/pkg/infra/repos"

	trace "github.com/BetaLixT/appInsightsTrace"
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	config.NewInsightsConfig,
	insights.NewInsights,
	db.GetClient,
	repos.NewTableStorageRepo,
	wire.Bind(
		new(tablestorage.ITableStorageRepository),
		new(*repos.TableStorageRepository),
	),
	config.NewDatabaseOptions,
	NewInfrastructure,
)

type Infrastructure struct {
	insightsCore  *trace.AppInsightsCore
	loggerFactory *logger.LoggerFactory
}

func NewInfrastructure(
	insightsCore *trace.AppInsightsCore,
	loggerFactory *logger.LoggerFactory,
) *Infrastructure {
	return &Infrastructure{
		insightsCore:  insightsCore,
		loggerFactory: loggerFactory,
	}
}

func (infra *Infrastructure) Start() {

}

func (infra *Infrastructure) Stop() {
	infra.insightsCore.Close()
	infra.loggerFactory.Close()
}
