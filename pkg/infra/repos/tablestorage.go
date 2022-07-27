package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"rabbitmqdddv2/pkg/domain/tablestorage"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"go.uber.org/zap"
)

type TableStorageRepository struct {
	client *aztables.Client
}

var _ tablestorage.ITableStorageRepository = (*TableStorageRepository)(nil)

func (repo *TableStorageRepository) SubTS(log *tablestorage.TableStorage, logger *zap.Logger, topic string) {
	repo.Table(logger)
	repo.Persist(log, logger, topic)
}

func (repo *TableStorageRepository) Persist(log *tablestorage.TableStorage, logger *zap.Logger, topic string) {
	count := repo.GetLength(logger) + 1
	log.RowKey = strconv.Itoa(count)
	marshalled, err := json.Marshal(log)
	if err != nil {
		panic(err)
	}
	_, err = repo.client.AddEntity(context.TODO(), marshalled, nil)
	if err != nil {
		panic(err)
	}
	// for err2 != nil { //pretty inefficient, find something quicker
	// 	count = count + 1
	// 	fmt.Println(count)

	// 	log.RowKey = strconv.Itoa(count)
	// 	marshalled, err := json.Marshal(log)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	_, err2 = repo.client.AddEntity(context.TODO(), marshalled, nil) // TODO: Check access policy, need Storage Table Data Contributor role
	// }
}

// func (repo *TableStorageRepository) Channel(logger *zap.Logger, topic string) {
// 	manager := usago.NewChannelManager("amqp://guest:guest@localhost:55001/", logger)
// 	bldr := usago.NewChannelBuilder().WithQueue(
// 		topic,
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	).WithConfirms(true)
// 	chnl, err := manager.NewChannel(*bldr)
// 	if err != nil {
// 		_, src, line, _ := runtime.Caller(0)
// 		repo.Log("", "", src, line, logger, "failed to create channel")
// 		fmt.Printf("failed to create channel")
// 		return
// 	}
// 	consumer, _ := chnl.RegisterConsumer(
// 		topic,
// 		"",
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	fmt.Println("CONSUMER REGISTERED")
// 	go func() {
// 		for {
// 			msg := <-consumer
// 			_, src, line, _ := runtime.Caller(0)
// 			repo.Log(string(msg.Body), msg.RoutingKey, src, line, logger, "message read")
// 		}

// 	}()

// }

// func (repo *TableStorageRepository) Log(body string, routingKey string, src string, line int, logger *zap.Logger, message string) {

// 	count := repo.GetLength(logger) + 1
// 	log.RowKey = strconv.Itoa(count)
// 	fmt.Println(log)
// 	marshalled, err := json.Marshal(log)
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = repo.client.AddEntity(context.TODO(), marshalled, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// for err2 != nil { //pretty inefficient, find something quicker
// 	// 	count = count + 1
// 	// 	fmt.Println(count)

// 	// 	log.RowKey = strconv.Itoa(count)
// 	// 	marshalled, err := json.Marshal(log)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	_, err2 = repo.client.AddEntity(context.TODO(), marshalled, nil) // TODO: Check access policy, need Storage Table Data Contributor role
// 	// }
// 	fmt.Println("DONE")

// }

func (repo *TableStorageRepository) Table(logger *zap.Logger) {
	_, err := repo.client.CreateTable(context.TODO(), nil)
	if err != nil {
		fmt.Println("Table Created")
	}
}

func (repo *TableStorageRepository) GetLength(logger *zap.Logger) int {
	listPager := repo.client.NewListEntitiesPager(nil)
	total_count := 0
	for listPager.More() {
		response, err := listPager.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}
		total_count += len(response.Entities)
	}
	return total_count

}

func Ptr[T any](v T) *T {
	return &v
}

func NewTableStorageRepo(client *aztables.Client) *TableStorageRepository {
	fmt.Println(client)
	return &TableStorageRepository{
		client: client,
	}
}
