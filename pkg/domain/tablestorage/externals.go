package tablestorage

import "go.uber.org/zap"

type ITableStorageRepository interface {
	SubTS(*TableStorage, *zap.Logger, string)
}
