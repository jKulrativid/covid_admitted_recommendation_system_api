package database

import (
	"os"

	"github.com/go-redis/redis/v7"
)

type RedisClient interface {
	Get()
	Set()
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

func (r *redisClient) Get() {

}

func (r *redisClient) Set() {

}
