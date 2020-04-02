package main

import (
	"01-helloworld/protos"
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
)

type Greeter struct {
}

func (g *Greeter) Hello(context context.Context, req *protos.Request, rsp *protos.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)
	service.Init()

	err := protos.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
