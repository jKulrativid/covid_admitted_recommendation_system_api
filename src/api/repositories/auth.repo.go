package repositories

import (
	"covid_admission_api/database"
)

type AuthRepo interface {
	GetFromClient(key string) (string, error)
}

type authRepo struct {
	redis database.RedisClient
}

func NewAuthRepo(rs database.RedisClient) AuthRepo {
	return &authRepo{
		redis: rs,
	}
}

func (a *authRepo) GetFromClient(key string) (string, error) {
	result, err := a.redis.Get(key)
	if err != nil {
		return "", err
	}
	return result, nil
}
