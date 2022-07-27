package db

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func GetClient() *aztables.Client {
	accountName := "gostorageaccount123" //export AZURE_STORAGE_ACCOUNT=gostorageaccount123   export AZURE_TABLE_NAME=gostoragetable
	tableName := "rabbitmqlogs"
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}
	serviceURL := fmt.Sprintf("https://%s.table.core.windows.net/%s", accountName, tableName)
	client, err := aztables.NewClient(serviceURL, cred, nil)
	if err != nil {
		panic(err)
	}
	return client
}
