package main

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	fmt.Println(carbon.Now().TimeString())
	fmt.Println(carbon.Now().DateTimeString())

	today, _ := carbon.NowInLocation("America/Vancouver")
	fmt.Println("美国时间", today)

	fmt.Println("明天的时间", carbon.Now().AddDay())
	fmt.Println("明年的时间", carbon.Now().AddYear())

	nextOlympics, _ := carbon.CreateFromDate(2016, time.August, 5, "Europe/London")
	fmt.Println("下一次奥林匹克运动会时间", nextOlympics.AddYears(4))

	diff := carbon.Now().DiffInDays(carbon.Now().SubDays(3), true)
	fmt.Println(diff)

	now, err := carbon.CreateFromFormat("2016-12-11", "2019-01-21 10:10:10", "Local")
	if err != nil {
		panic(err)
	}
	fmt.Println(now)
}
