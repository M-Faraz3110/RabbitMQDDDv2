package logger

import (

	// "time"

	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

type LoggerFactory struct {
	lgr *zap.Logger
}

func NewLoggerFactory() (*LoggerFactory, error) {
	lgr, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &LoggerFactory{
		lgr: lgr,
	}, nil
}

func (lf *LoggerFactory) NewLogger() *zap.Logger {
	return lf.lgr
}

func (lf *LoggerFactory) Close() {
	lf.lgr.Sync()
}
