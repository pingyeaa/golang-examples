package main

import (
	"errors"
	"flag"
	"fmt"
)

type user string

func (u *user) String() string {
	return fmt.Sprint(*u)
}

func (u *user) Set(value string) error {
	if len(value) < 3 {
		return errors.New("姓名长度不得小于3位")
	}
	*u = user(value + "_suffix")
	return nil
}

var userFlag user

func main() {
	flag.Var(&userFlag, "user", "用户名")
	flag.Parse()
	fmt.Println(userFlag)
}
