package database

import (
	"os"
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisClient interface {
	Get(key string) (string, error)
	Set(key, val string, exp time.Duration) error
	Del(key string) error
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient() (RedisClient, error) {
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
	return &redisClient{client}, nil

}

func (r *redisClient) Get(key string) (string, error) {
	result := r.client.Get(key)
	return result.String(), result.Err()
}

func (r *redisClient) Set(key, val string, exp time.Duration) error {
	return r.client.Set(key, val, exp).Err()
}

func (r *redisClient) Del(key string) error {
	return r.client.Del(key).Err()
}
