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
	//  lvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	// 	return true
	// })

	// ws, closeOut, err := zap.Open("stderr")
	// if err != nil {
	// 	return nil, err
	// }
	// errws, _, err := zap.Open("stderr")
	// if err != nil {
	// 	closeOut()
	// 	return nil, err
	// }
	//
	// ins := zapcore.WriteSyncer(&zapcore.BufferedWriteSyncer{
	//   WS: ws,
	// })
	// enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//
	// core := zapcore.NewTee(
	// 	zapcore.NewCore(enc, ins, lvl),
	// )
	//
	// logger := zap.New(
	//   core,
	//   zap.ErrorOutput(errws),
	//   zap.AddStacktrace(zapcore.ErrorLevel),
	//   zap.WrapCore(func(core zapcore.Core) zapcore.Core {
	// 		var samplerOpts []zapcore.SamplerOption
	// 		return zapcore.NewSamplerWithOptions(
	// 			core,
	// 			time.Second,
	// 			100,
	// 			100,
	// 			samplerOpts...,
	// 		)
	// 	}),
	// )
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
