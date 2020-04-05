package main

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

type A uint

//func getColors() Color {
//
//}

func main() {
	fmt.Println("Hello,", Magenta("Aurora"))
	fmt.Println(Bold(Cyan("Cya!")))

	msg := fmt.Sprintf("My name is %s", Green("pingyeaa"))
	fmt.Println(msg)

	//链式调用
	fmt.Println(BgGreen(Bold(Red("pingyeaa"))))
	fmt.Println(Red("haha").Bold().BgGreen())

	//fmt.Println(Colorize("Greeting", getColors()))
	fmt.Println(Colorize("Greeting", GreenFg|BoldFm|BrightBg))

	fmt.Println(Red("x").Colorize(GreenFg))

	fmt.Println(Gray(1-1, " 00-23 ").BgGray(24 - 1))

	//灰阶
	fmt.Println("  ",
		Gray(1-1, " 00-23 ").BgGray(24-1),
		Gray(4-1, " 03-19 ").BgGray(20-1),
		Gray(8-1, " 07-15 ").BgGray(16-1),
		Gray(12-1, " 11-11 ").BgGray(12-1),
		Gray(16-1, " 15-07 ").BgGray(8-1),
		Gray(20-1, " 19-03 ").BgGray(4-1),
		Gray(24-1, " 23-00 ").BgGray(1-1),
	)

	fmt.Println(Blink("Blink"))

	//%T
	r := Red("red")
	var i int
	fmt.Printf("%T %p\n", r, Green(&i))
	fmt.Println(Red(fmt.Sprintf("%T %p\n", r, Green(&i))))

}
