package databases

import (
	"os"

	"github.com/go-redis/redis/v7"
)

var RedisClient *redis.Client

func NewRedisClient() (*redis.Client, error) {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err

	}
	return client, nil

}
