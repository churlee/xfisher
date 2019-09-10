package config

import (
	"github.com/go-redis/redis"
	"log"
)

var (
	redisClient *redis.Client
)

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
	}
	redisClient = client
}

func GetRedisClient() *redis.Client {
	return redisClient
}
