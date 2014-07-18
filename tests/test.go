package main

import (
	"fmt"
	"github.com/liulixiang1988/WIM-System/models/datalineinfo"
	"time"
)

func main() {
	last, err := datalineinfo.Last()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(last.Id)

	day, _ := time.Parse("2006-01-02", "2014-07-09")
	results, _ := datalineinfo.WorkShift(day, 1)
	fmt.Println("result长度", len(results))
}
