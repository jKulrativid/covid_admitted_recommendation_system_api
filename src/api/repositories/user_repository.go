package repositories

import (
	"covid_admission_api/entities"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	conn   *gorm.DB
	client *redis.Client
}

func NewUserRepository(cn *gorm.DB, rc *redis.Client) *UserRepository {
	return &UserRepository{
		conn:   cn,
		client: rc,
	}
}

func (repo *UserRepository) Register(newUser *entities.User) error {
	return nil
}

func (repo *UserRepository) AddTokenToClient(token string, strID string, exp time.Duration) error {
	return repo.client.Set(token, strID, exp).Err()

}
