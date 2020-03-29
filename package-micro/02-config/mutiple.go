package main

import (
	"fmt"

	"github.com/micro/go-micro/v2/config/source/file"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/env"
)

func main() {
	err := config.Load(
		env.NewSource(),
		file.NewSource(
			file.WithPath("config.json"),
		),
	)
	if err != nil {
		panic(err)
	}
	conf := config.Map()
	fmt.Println(conf)
}
