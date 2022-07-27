package v1

import (
	"rabbitmqdddv2/pkg/domain/rabbitmq"
	"rabbitmqdddv2/pkg/infra/logger"

	"go.uber.org/zap"
)

type RabbitMqController struct {
	svc    *rabbitmq.RabbitMqService
	logger *zap.Logger
	// channel *usago.ChannelContext
}

// @BasePath

// ListForecasts godoc
// @Summary List forecasts
// @Schemes
// @Description List forecasts
// @Tags role
// @Accept json
// @Produce json
// @Success 200 {object} []res.ForecastDetailed{}
// @Router /api/v1/forecasts/ [get]
func (ctrl *RabbitMqController) Csub(logger *zap.Logger, topic string) {

	ctrl.svc.SubService(topic, logger)

}

func NewRabbitMqController(
	svc *rabbitmq.RabbitMqService,
	lf *logger.LoggerFactory,
) *RabbitMqController {

	return &RabbitMqController{
		svc:    svc,
		logger: lf.NewLogger(),
	}
}
