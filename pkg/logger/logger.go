package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	dev  = "dev"
	prod = "prod"
)

func NewLogger(env string) *zap.Logger {
	var (
		logger *zap.Logger
		err    error
		config zap.Config
	)

	switch env {
	case dev:
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = config.Build()
	case prod:
		config = zap.NewProductionConfig()
		logger, err = config.Build()
	default:
		log.Fatal("Logger is uninitialized: wrong env value")
	}

	if err != nil {
		log.Fatal("error configuring logger: " + err.Error())
	}

	return logger

}
