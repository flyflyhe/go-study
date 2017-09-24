package main

import(
	"fmt"
	"services"
)

func main() {
	redisClient := services.GetRedis()
	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)

	redisClient.HSet("h", "a", "1")
	redisClient.HSet("h", "b", "2")
	tmp,_ := redisClient.HGetAll("h").Result()
	fmt.Println(tmp["a"])
}