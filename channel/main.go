package main

import "fmt"

func main() {

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 2

	close(ch)

	for d := range ch {
		fmt.Println(d)
	}
}
