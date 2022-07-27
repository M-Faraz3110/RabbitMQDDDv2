package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"rabbitmqdddv2/pkg/domain/rabbitmq"
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
	// repo.Table(logger)
	count := repo.GetLength(logger) + 1
	log.RowKey = strconv.Itoa(count)
	marshalled, err := json.Marshal(log)
	if err != nil {
		log.Body = "Could not marshal log"
		rabbitmq.Log(log.Msg, log.Body, topic, logger)
	}
	_, err = repo.client.AddEntity(context.TODO(), marshalled, nil)
	if err != nil {
		log.Body = "Could not add item to storage table"
		rabbitmq.Log(log.Msg, log.Body, topic, logger)
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

func NewTableStorageRepo(newclient *aztables.Client) *TableStorageRepository {
	_, err := newclient.CreateTable(context.TODO(), nil)
	if err != nil {

	} else {
		fmt.Println("Creating Table")
	}
	return &TableStorageRepository{
		client: newclient,
	}
}
