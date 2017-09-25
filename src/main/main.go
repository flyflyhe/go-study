package main

import (
	"fmt"
	"services"
)

func main() {
	redis := services.GetRedis()

	pong, err := redis.Ping().Result()
	fmt.Println(pong, err)
}
