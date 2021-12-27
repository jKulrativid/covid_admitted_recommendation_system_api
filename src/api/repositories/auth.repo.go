package repositories

import "covid_admission_api/database"

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
	return a.redis.Get(key)
}
