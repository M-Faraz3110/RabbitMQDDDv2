package tablestorage

import "go.uber.org/zap"

type ITableStorageRepository interface {
	SubTS(*zap.Logger, string)
}
