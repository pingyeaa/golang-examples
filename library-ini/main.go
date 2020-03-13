package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("文件读取错误", err)
		os.Exit(1)
	}
	fmt.Println(cfg.Section("").Key("username"))
	fmt.Println(cfg.Section("mysql").Key("username"))

	//限制值
	fmt.Println(cfg.Section("mysql").Key("username").In("张三", []string{"张三", "李四"}))

	//修改配置文件
	cfg.Section("mysql").Key("username").SetValue("李四")
	err = cfg.SaveTo("config.ini")
	if err != nil {
		fmt.Println("文件保存错误", err)
	}
}
