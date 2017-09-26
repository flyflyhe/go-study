package redis

import (
	"config"
	"github.com/go-redis/redis"
)

func GetRedis5() *redis.Client {
	config := config.GetRedis()

	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       6,
	})

	return client
}

func GetRedis6() *redis.Client {
	config := config.GetRedis()

	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       7,
	})

	return client
}
