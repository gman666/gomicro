package main

import (
	"fmt"
	"time"
	micro "github.com/micro/go-micro"
)

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[%v] server request: %s", time.Now(), req.Endpoint())
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