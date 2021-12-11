package repositories

import (
	"log"
	"os"

	"github.com/go-redis/redis/v7"
)

type JWTTokenRepository struct {
	client *redis.Client
}

func NewJWTTokenRepository() *JWTTokenRepository {
	redisClient := initClient()
	return &JWTTokenRepository{
		client: redisClient,
	}
}

func initClient() *redis.Client {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	if _, err := client.Ping().Result(); err != nil {
		log.Fatal("Crashed in JWTinitClient (jwt_token_repository.go) : Could Not Connect To Redis at ", dsn)
	}
	return client

}
