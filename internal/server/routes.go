package server

import (
	"context"
	"e-commerce-api/internal/constants"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerAbs struct {
}
type AppConfig struct {
	Port        string        `yaml:"serv_port" env-default:"8080"`
	SecretJwt   string        `yaml:"secret_jwt" env-required:"true"`
	RWTimeout   time.Duration `yaml:"rw_timeout" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"120s"`
}

type Router struct {
	handler HandlerAbs
	logger  *zap.Logger
	config  AppConfig
	engine  *gin.Engine
}

func NewRouter(logger *zap.Logger, config AppConfig, handler HandlerAbs, engine *gin.Engine) *Router {

	return &Router{
		handler: handler,
		logger:  logger,
		config:  config,
		engine:  engine,
	}
}

func (r *Router) Run() {
	srv := r.engine
	// nado config
	srv.Run(":8080")

	srv.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

}

func (r *Router) Shutdown(fn func(ctx context.Context)) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), constants.ShutdownTime)
	defer cancel()

	fn(ctx)
}
