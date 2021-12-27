package services

import (
	"covid_admission_api/database"
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

/* Register Test */
func TestRegister(t *testing.T) {
	db, _, err := database.NewMockDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	rs, err := database.NewMockRedisClient()
	if err != nil {
		t.Fatal(err)
	}
	userRepo := repositories.NewUserRepository(db, rs)
	userService := NewUserService(userRepo)

	t.Run("normal register", func(t *testing.T) {
		var err error
		user := &entities.UserRegister{
			UserName: "Foo Bar",
			Email:    "foobar@gmail.com",
			Password: "123456789",
		}
		err = userService.Register(user)
		assert.NoError(t, err)
	})
	t.Run("duplicate register", func(t *testing.T) {
		var err error
		user := &entities.UserRegister{
			UserName: "Zinovacz Zell",
			Email:    "vaccine@gmail.com",
			Password: "123456789",
		}
		err = userService.Register(user)
		assert.Error(t, err)
		//err = userService.Register(user)
		//if assert.Error()
	})

}

/* Sign In Test */
