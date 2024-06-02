package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/log"

	component "github.com/josuedeavila/supreme-palm-tree/internal"
)

const (
	shutdownTimeout = 5 * time.Second
	apiTimeout      = 10 * time.Second
)

var (
	httpPort = 8081
)

func main() {
	ctx := context.Background()

	// component
	c := component.New(nil, nil)

	// create engine
	engine := gin.New()

	// add middlewares
	engine.Use(gin.Recovery())

	routes := engine.Group("/")
	// register routes
	c.RegisterBalanceRoutes(routes)
	c.RegisterEventRoutes(routes)
	c.RegisterResetRoutes(routes)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", httpPort),
		Handler:      engine,
		ReadTimeout:  apiTimeout,
		WriteTimeout: apiTimeout,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("failed to serve http server with error: %v", err)
		}
	}()
	log.Info("Server started")

	<-done
	log.Info("Server stopped")

	srvCtx, sCancel := context.WithTimeout(ctx, shutdownTimeout)
	defer sCancel()

	if err := srv.Shutdown(srvCtx); err != nil {
		log.Panicf("Server shutdown Failed:%+v", err)
	}
	log.Info("Server exited properly")
}
