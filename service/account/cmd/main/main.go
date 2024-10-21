package main

import (
	"account/internal/config"
	"account/internal/http/server"
	"account/pkg/logger"
	"context"
)

const (
	serviceName = "account"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New(serviceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)
	cfg := config.New()
	if cfg == nil {
		panic("failed to load config")
	}

	//db, err := postgres.New(cfg.Config)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}

	echoServer := server.New(cfg.HTTPServerPort)
	if err := echoServer.Start(); err != nil {
		mainLogger.Error(ctx, err.Error())
	}
}
