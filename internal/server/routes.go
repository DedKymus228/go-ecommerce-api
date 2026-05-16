package server

import (
	"e-commerce-api/internal/infrastructure/config"
	"e-commerce-api/internal/server/handlers"
	"e-commerce-api/internal/server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	handler handlers.Handler
	logger  *zap.Logger
	config  config.AppConfig
	engine  *gin.Engine
	srv     *http.Server
	md      *middleware.Middleware
}

func NewRouter(logger *zap.Logger, config config.AppConfig, handler *handlers.Handler, md *middleware.Middleware) *Router {
	engine := gin.Default()

	return &Router{
		handler: *handler,
		logger:  logger,
		config:  config,
		engine:  engine,
		md:      md,
	}
}
