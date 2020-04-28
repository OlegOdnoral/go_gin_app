package redis_adapter

import (
	"github.com/go-redis/redis"
	"time"
)

var (
	redisClient *redis.Client
)

func onConnectHook(*redis.Conn) error {

}

func InitRedisClient() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "",
		OnConnect:  onConnectHook,
		MaxRetries: 3,
		MaxConnAge: time.Hour * 3,
		//DB:       0,
	})
	redisClient.Close()
}
