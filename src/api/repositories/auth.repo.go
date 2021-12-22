package repositories

import "covid_admission_api/database"

type AuthRepo interface {
	GetFromClient(key string) (string, error)
}

type authRepo struct {
	client database.RedisClient
}

func NewAuthRepo(rs database.RedisClient) AuthRepo {
	return &authRepo{
		client: rs,
	}
}

func (a *authRepo) GetFromClient(key string) (string, error) {
	return a.client.Get(key)
}
