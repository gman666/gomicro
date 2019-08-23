package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	proto "github.com/gman666/gomicro/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter.client"),
	)

	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Gordon"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Greeting)

	rsp2, err := greeter.Goodbye(context.TODO(), &proto.GoodbyeRequest{Name: "Gordon"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp2.Greeting)
}