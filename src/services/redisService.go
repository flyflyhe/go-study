package services

import (
	"config"
	"fmt"
	"github.com/go-redis/redis"
)

func GetRedis() *redis.Client {
	config := config.GetRedis()

	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.Select,
	})

	return client
}
