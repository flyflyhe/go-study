package main

import (
	"fmt"
	"services"
)

func main() {
	data := services.GetUserSpeedHistoryData()
	if data == nil {
		fmt.Println("异常")
	} else {
		fmt.Println("结束")
	}
	fmt.Println(data)
}
