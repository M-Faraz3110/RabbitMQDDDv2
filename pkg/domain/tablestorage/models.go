package tablestorage

import "github.com/Azure/azure-sdk-for-go/sdk/data/aztables"

type TableStorage struct { //CHANGE ACCORDING TO UDC
	aztables.Entity

	Level      string
	Ts         string
	Caller     string
	Msg        string
	Body       string
	RoutingKey string
}
