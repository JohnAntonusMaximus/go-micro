package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
