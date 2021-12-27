package repositories

import (
	"covid_admission_api/database"
	"covid_admission_api/entities"
	"time"
)

type UserRepository interface {
	CreateNewUser(newUser *entities.User) error
	GetUserFromUserName(user *entities.User, userName string) error
	GetUserFromEmail(user *entities.User, email string) error
	SaveJWT(key, val string, exp time.Duration) error
	DeleteJWT(key string) error
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

func (u *userRepository) CreateNewUser(newUser *entities.User) error {
	db := u.database.GetConnection()
	if err := db.Create(newUser).Error; err != nil {
		return err
		// return entities.ErrorConflict
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

func (u *userRepository) GetUserFromEmail(user *entities.User, email string) error {
	db := u.database.GetConnection()
	if err := db.Find("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) SaveJWT(key, val string, exp time.Duration) error {
	return u.redis.Set(key, val, exp)
}

func (u *userRepository) DeleteJWT(key string) error {
	return u.redis.Del(key)
}
