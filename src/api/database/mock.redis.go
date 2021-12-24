package database

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v7"
)

func NewMockRedisClient() (RedisClient, error) {
	mockRedis, err := miniredis.Run()
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(&redis.Options{
		Addr: mockRedis.Addr(),
	})
	return &redisClient{client}, nil
}
