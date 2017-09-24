package services

import(
	"github.com/go-redis/redis"
)

func GetRedis() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	return client
}