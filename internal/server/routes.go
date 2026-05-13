package server

import (
	"e-commerce-api/internal/infrastructure/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerAbs struct {
}
type Router struct {
	handler HandlerAbs
	logger  *zap.Logger
	config  config.AppConfig
	engine  *gin.Engine
	srv     *http.Server
}

func NewRouter(logger *zap.Logger, config config.AppConfig, handler HandlerAbs) *Router {
	engine := gin.Default()

	return &Router{
		handler: handler,
		logger:  logger,
		config:  config,
		engine:  engine,
	}
}
