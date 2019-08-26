package main

import (
	"context"
	"fmt"
	"log"
	"github.com/micro/go-micro/server"
	micro "github.com/micro/go-micro"
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("service: %s, request: %v", req.Endpoint(), req.Body())
		return fn(ctx, req, rsp)
	}
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.WrapHandler(logWrapper),
	)

	service.Init()
	service.Run()
}