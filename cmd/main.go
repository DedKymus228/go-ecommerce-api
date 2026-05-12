package main

import (
	"e-commerce-api/internal/infrastructure/config"
	mydb "e-commerce-api/internal/repository/db/sqlc"
	"e-commerce-api/internal/server"
	logging "e-commerce-api/pkg/logger"
	"e-commerce-api/pkg/postgre"
	"log"

	"go.uber.org/zap"
)

var configPath = "config/config.yaml"

func main() {
	cfg := config.GetConfig(configPath)
	logger := logging.NewLogger(cfg.Env)

	logger.Info("Logger is set up",
		zap.String("env:", cfg.Env))

	dbpool, err := postgre.NewPgxPool(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbpool.Close()

	db := mydb.New(dbpool)
	_ = db

	logger.Info("Database connection established")

	if err = postgre.RunMigrations(cfg.DB); err != nil {
		logger.Fatal("migrations fault", zap.Error(err))
	}

	logger.Info("Migrations applied successfully")

	server.Run()

	logger.Info("Server is running")

}

//  serv
