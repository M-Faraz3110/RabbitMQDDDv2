package config

import (
	"rabbitmqdddv2/pkg/infra/db"

	trace "github.com/BetaLixT/appInsightsTrace"
)

// TODO: Implement something like viper here
func NewInsightsConfig() *trace.AppInsightsOptions {
	return &trace.AppInsightsOptions{
		ServiceName: "EventLogger",
	}
}

func NewDatabaseOptions() *db.Options {
	return &db.Options{
		DatabaseServiceName: "main-database",
	}
}
