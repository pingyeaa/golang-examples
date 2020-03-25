package main

import (
	"flag"
	"fmt"
)

var port int

func init() {
	flag.IntVar(&port, "port", 80, "端口号")
}

func main() {
	var ip = flag.String("ip", "127.0.0.1", "此处传入IP地址")

	flag.Parse()
	fmt.Println("ip", *ip)
	fmt.Println("port", port)
}
