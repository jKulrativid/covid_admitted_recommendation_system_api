package repositories

import (
	"covid_admission_api/entities"
	"log"
	"os"

	"github.com/go-redis/redis/v7"
)

type UserRepository struct {
	// implements gorm here
	redisClient *redis.Client
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		redisClient: newRedisClient(),
	}
}

func (repo *UserRepository) Register(newUser *entities.User) error {
	return nil
}

func newRedisClient() *redis.Client {
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
