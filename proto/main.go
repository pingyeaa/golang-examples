package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	test := &Test{
		Label: "a",
		Type:  32,
		Reps:  []int64{10, 11},
	}
	resp, err := proto.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
