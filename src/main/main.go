package main

import (
	"fmt"
	"services"
	sj "services/json"
)

func main() {
	data := services.GetUserSpeedHistoryData()
	if data == nil {
		fmt.Println("异常")
	} else {
		fmt.Println("结束")
	}
	fmt.Println(data)
	game := sj.GetGame()
	fmt.Println(game)
}
