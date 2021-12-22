package repositories

import (
	"covid_admission_api/database"
	"covid_admission_api/entities"
	"time"
)

type UserRepository interface {
	RegisterNewUser(newUser *entities.User) error
	GetUserFromUserName(user *entities.User, userName string) error
	SetToRedis(key, val string, exp time.Duration) error
}

type userRepository struct {
	database database.Database
	redis    database.RedisClient
}

func NewUserRepository(db database.Database, rs database.RedisClient) UserRepository {
	return &userRepository{
		database: db,
		redis:    rs,
	}
}

func (u *userRepository) RegisterNewUser(newUser *entities.User) error {
	db := u.database.GetConnection()
	if err := db.Create(newUser).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUserFromUserName(user *entities.User, userName string) error {
	db := u.database.GetConnection()
	if err := db.Where("user_name = ?", userName).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) SetToRedis(key, val string, exp time.Duration) error {
	return u.redis.Set(key, val, exp)
}
