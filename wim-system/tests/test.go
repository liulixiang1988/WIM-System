package main

import (
	"fmt"
	"github.com/liulixiang1988/WIM-System/wim-system/models/datalineinfo"
)

func main() {
	last, err := datalineinfo.Last()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(last.Id)

}
