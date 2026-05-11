package main

import (
	"e-commerce-api/internal/infrastructure/config"
	logging "e-commerce-api/pkg/logger"

	"go.uber.org/zap"
)

var configPath = "config/config.yaml"

func main() {
	cfg := config.GetConfig(configPath)
	logger := logging.NewLogger(cfg.Env)

	logger.Info("Logger is set up",
		zap.String("env:", cfg.Env))
}

// loggin db serv
