package repositories

import "covid_admission_api/entities"

type UserRepository struct {
	// implements gorm here
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) Register(newUser *entities.User) error {
	return nil
}
