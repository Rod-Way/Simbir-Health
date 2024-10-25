package main

import (
	"account/internal/config"
	"account/internal/http/server"
	"context"
	"log"
)

const (
	serviceName = "account"
)

func main() {
	ctx := context.Background()
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
		log.Fatal(ctx, err.Error())
	}
}
