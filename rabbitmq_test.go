package rabbitmqtest

import (
	"fmt"
	v1 "rabbitmqdddv2/pkg/app/controllers/v1"
	"rabbitmqdddv2/pkg/domain/rabbitmq"
	"rabbitmqdddv2/pkg/infra/db"
	"rabbitmqdddv2/pkg/infra/logger"
	"rabbitmqdddv2/pkg/infra/repos"
	"strconv"
	"testing"
	"time"

	"github.com/BetaLixT/usago"
	amqp "github.com/rabbitmq/amqp091-go"
)

func TestNewChannelManager(t *testing.T) {
	lf, err := logger.NewLoggerFactory()
	manager := usago.NewChannelManager("<RABBITMQ_CONNECTION_STRING>", lf.NewLogger())
	bldr := usago.NewChannelBuilder().WithQueue(
		"Notification",
		false,
		false,
		false,
		false,
		nil,
	).WithConfirms(true)
	chnl, err := manager.NewChannel(*bldr)
	if err != nil {
		fmt.Printf("failed to create channel")
		return
	}
	// consume
	client := db.GetClient()
	repo := repos.NewTableStorageRepo(client)
	svc := rabbitmq.NewRabbitMqService(repo)
	v1.NewRabbitMqController(svc, lf).Csub(lf.NewLogger(), "Notification")
	i := 0
	for i < 10 {
		body := "testmf" + strconv.Itoa(i)
		_, err = chnl.Publish(
			"",
			"Notification",
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		for err != nil {
			_, err = chnl.Publish(
				"",
				"Notification",
				true,  // mandatory
				false, // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				},
			)
		}
		i += 1
		time.Sleep(1 * time.Second)
	}
}
