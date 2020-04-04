package main

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Println("Hello,", Magenta("Aurora"))
	fmt.Println(Bold(Cyan("Cya!")))
}
