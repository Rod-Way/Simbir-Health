package main

import (
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
}
