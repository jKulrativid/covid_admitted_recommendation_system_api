package repositories

import (
	"covid_admission_api/entities"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	conn        *gorm.DB
	redisClient *redis.Client
}

func NewUserRepository(cn *gorm.DB, rc *redis.Client) *UserRepository {
	return &UserRepository{
		conn:        cn,
		redisClient: rc,
	}
}

func (repo *UserRepository) Register(newUser *entities.User) error {
	return nil
}
