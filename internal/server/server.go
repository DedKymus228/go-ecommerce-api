package server

import (
	"context"
	"e-commerce-api/internal/constants"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func (r *Router) Shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), constants.ShutdownTime)
	defer cancel()

	if err := r.srv.Shutdown(ctx); err != nil {
		r.logger.Fatal("Server forced to shutdown: ", zap.Error(err))
	}

	r.logger.Info("Server exiting")

}

func (r *Router) Run() {
	r.Map()
	r.srv = &http.Server{
		Addr:         ":" + r.config.Port,
		Handler:      r.engine,
		ReadTimeout:  r.config.RWTimeout,
		WriteTimeout: r.config.RWTimeout,
		IdleTimeout:  r.config.IdleTimeout,
	}

	go func() {
		if err := r.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			r.logger.Fatal("Server start failed", zap.Error(err))

		}
	}()

}
