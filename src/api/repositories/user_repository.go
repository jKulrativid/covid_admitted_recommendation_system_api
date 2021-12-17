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

func (repo *UserRepository) Register(newUser *entities.User) (err error) {
	if err = repo.conn.Create(newUser).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) PullUserData(user *entities.User, userName string) (err error) {
	if err = repo.conn.Where("user_name = ?", userName).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) AddTokenToClient(token string, strID string, exp time.Duration) error {
	return repo.client.Set(token, strID, exp).Err()

}

func (repo *UserRepository) GetFromClient(accessUuid string) (string, error) {
	return repo.client.Get(accessUuid).Result()
}
