package redis

import (
	"github.com/go-redis/redis"
	"log"
	"fmt"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	result, err :=Redis.Ping().Result()

	if err != nil {
		log.Fatal(fmt.Errorf("redis NewClient Ping failed: %v result:%s", err, result))
		return
	}
	fmt.Println("Connnect Redis Success")
}