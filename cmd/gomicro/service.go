package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	proto "github.com/gman666/gomicro/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hey there " + req.Name
	return nil
}

func (g *Greeter) Goodbye(ctx context.Context, req *proto.GoodbyeRequest, rsp *proto.GoodbyeResponse) error {
	rsp.Greeting = "Goodbye " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// parse command line flags
	service.Init()

	// register a handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
