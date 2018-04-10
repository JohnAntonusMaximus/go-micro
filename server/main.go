package main

import (
	"context"
	"fmt"

	proto "github.com/johnantonusmaximus/go-micro/proto"
	micro "github.com/micro/go-micro"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello" + req.Name

	fmt.Printf("Responding with %s\n", rsp.Greeting)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(proto.Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
