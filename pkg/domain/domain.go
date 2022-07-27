package domain

import (
	"rabbitmqdddv2/pkg/domain/rabbitmq"

	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	rabbitmq.NewRabbitMqService,
)
