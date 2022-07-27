package rabbitmq

import (
	"fmt"
	"rabbitmqdddv2/pkg/domain/tablestorage"
	"runtime"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/BetaLixT/usago"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type RabbitMqService struct {
	repo tablestorage.ITableStorageRepository
}

func (svc *RabbitMqService) SubService(
	topic string,
	logger *zap.Logger,
) {
	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55001/", logger)
	bldr := usago.NewChannelBuilder().WithQueue(
		topic,
		false,
		false,
		false,
		false,
		nil,
	).WithConfirms(true)
	chnl, err := manager.NewChannel(*bldr)
	if err != nil {
		// _, src, line, _ := runtime.Caller(0)
		// repo.Log("", "", src, line, logger, "failed to create channel")
		fmt.Printf("failed to create channel")
		return
	}
	consumer, _ := chnl.RegisterConsumer(
		topic,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	go func() {
		for {
			msg := <-consumer
			log := svc.Log("message read", string(msg.Body), topic, logger)
			svc.repo.SubTS(log, logger, topic)
			// repo.Log(string(msg.Body), msg.RoutingKey, src, line, logger, "message read")
		}
	}()

}

func (svc *RabbitMqService) Log(
	message string,
	body string,
	topic string,
	logger *zap.Logger,
) *tablestorage.TableStorage {
	_, src, line, _ := runtime.Caller(0)
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore)
	observedLogger.Info(
		message,
		zap.String("body", string(body)),
	)
	logger.Info(message,
		zap.String("body", string(body)))
	logbody := observedLogs.All()[0]
	log := tablestorage.TableStorage{
		Entity: aztables.Entity{
			PartitionKey: "1",
			RowKey:       "1",
		},
		Level:      logbody.Level.String(),
		Ts:         strconv.Itoa(int(logbody.Time.Unix())),
		Caller:     src + ":" + strconv.Itoa(line),
		Msg:        logbody.Message,
		Body:       string(body),
		RoutingKey: string(topic),
	}
	return &log
}

func NewRabbitMqService(
	repo tablestorage.ITableStorageRepository,
) *RabbitMqService {
	return &RabbitMqService{
		repo: repo,
	}
}
