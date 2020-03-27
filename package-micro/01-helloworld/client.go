package main

import (
	"01-helloworld/protos"
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	greeter := protos.NewGreeterService("greeter", service.Client())
	rsp, err := greeter.Hello(context.TODO(), &protos.Request{Name: "pingye"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Greeting)
}
