package main

import (
	"fmt"

	"github.com/micro/go-micro/v2/config"
)

type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func main() {
	err := config.LoadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	conf := config.Map()
	fmt.Println(conf["hosts"])

	//获取多级配置
	value := config.Get("hosts", "database", "address").String("abc")
	fmt.Println(value)

	//监控配置变动
	WatchPath()
}

func WatchPath() {
	w, err := config.Watch("hosts", "database")
	if err != nil {
		fmt.Println(err)
	}

	v, err := w.Next()
	if err != nil {
		fmt.Println(err)
	}

	var host Host
	err = v.Scan(&host)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(host)
}
