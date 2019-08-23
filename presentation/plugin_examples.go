package main

import (
	"github.com/micro/go-micro" 
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/transport/nats"
	"github.com/micro/go-plugins/broker/kafka"
)

func main() {
	registry := etcdv3.NewRegistry()
	broker := kafka.NewBroker()
	transport := nats.NewTransport()

	service := micro.NewService(
			micro.Name("greeter"),
			micro.Registry(registry),
			micro.Broker(broker),
			micro.Transport(transport),
	)

	service.Init()
	service.Run()
}
