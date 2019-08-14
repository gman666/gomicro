package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	proto "github.com/micro/examples/service/proto"
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
}