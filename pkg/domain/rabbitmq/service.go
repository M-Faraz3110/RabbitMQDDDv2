package rabbitmq

import (
	"rabbitmqdddv2/pkg/domain/tablestorage"

	"go.uber.org/zap"
)

type RabbitMqService struct {
	repo tablestorage.ITableStorageRepository
}

func (svc *RabbitMqService) SubService(
	topic string,
	logger *zap.Logger,
) {
	svc.repo.SubTS(logger, topic)
}

func NewRabbitMqService(
	repo tablestorage.ITableStorageRepository,
) *RabbitMqService {
	return &RabbitMqService{
		repo: repo,
	}
}
