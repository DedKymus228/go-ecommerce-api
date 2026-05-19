package main

import (
	"e-commerce-api/internal/infrastructure/auth"
	"e-commerce-api/internal/infrastructure/config"
	db "e-commerce-api/internal/repository/db/sqlc"
	"e-commerce-api/internal/server"
	"e-commerce-api/internal/server/handlers"
	"e-commerce-api/internal/server/middleware"
	"e-commerce-api/internal/service"
	logging "e-commerce-api/pkg/logger"
	"e-commerce-api/pkg/postgre"
	"log"

	"go.uber.org/zap"
)

var configPath = "config/config.yaml"

func main() {
	cfg := config.GetConfig(configPath)
	logger := logging.NewLogger(cfg.Env)
	tokenManager := auth.NewJWTManager(cfg.App.SecretJwt, cfg.App.TokenTTl)

	logger.Info("Logger is set up",
		zap.String("env:", cfg.Env))

	dbpool, err := postgre.NewPgxPool(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbpool.Close()

	logger.Info("Database connection established")

	repo := db.New(dbpool)

	// SERVICE
	orderService := service.NewOrderService(repo)
	authService := service.NewAuthService(repo, tokenManager, cfg.App.TokenTTl)
	productService := service.NewProductService(repo)
	cartService := service.NewCartService(repo)
	//

	// MIDDLEWARE
	md := middleware.NewMiddleware(tokenManager)
	//

	handler := handlers.NewHandler(orderService, cartService, authService, productService, logger)
	router := server.NewRouter(logger, cfg.App, handler, md)

	if err = postgre.RunMigrations(cfg.DB); err != nil {
		logger.Fatal("migrations fault", zap.Error(err))
	}
	logger.Info("Migrations applied successfully")

	logger.Info("Server is starting...")
	router.Run()

	router.Shutdown()
	// fix cart handler
	//TODO orders and Admin panel ,product service and handler

}
