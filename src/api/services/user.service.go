package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(newUser *entities.UserRegister) (handleError error)
	SignIn(user *entities.UserSignIn) (uuid string, err error)
	SignOut(user *entities.User) (err error)
	GenerateToken(userUuid string) (td *TokenDetail, err error)
	CreateAuth(uuid string, td *TokenDetail) (err error)
}

type userService struct {
	repo     repositories.UserRepository
	atSecret string
	rtSecret string
}

func NewUserService(r repositories.UserRepository) UserService {
	ass := os.Getenv("ACCESS_JWT_SECRET")
	rfs := os.Getenv("REFRESH_JWT_SECRET")
	if ass == "" || rfs == "" {
		log.Fatal("Crashed in NewJWTService (jwt_service.go) : No Environment Variable \"ACCESS_JWT_SECRET\" or \"REFRESH_JWT_SECRET\" Given")
	}
	return &userService{
		repo: r,
	}
}

func (u *userService) Register(newUser *entities.UserRegister) (handleError error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	user := &entities.User{
		Uuid:           uuid.NewV4().String(),
		UserName:       newUser.UserName,
		Email:          newUser.Email,
		HashedPassword: string(hashedPassword),
	}
	handleError = u.repo.RegisterNewUser(user)
	return handleError

}

func (u *userService) SignIn(userSignIn *entities.UserSignIn) (uuid string, err error) {
	var user entities.User
	err = u.repo.GetUserFromUserName(&user, userSignIn.UserName)
	if err != nil {
		return "", fmt.Errorf("username not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userSignIn.Password))
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}
	return uuid, nil
}

func (u *userService) SignOut(user *entities.User) (err error) {
	err = nil
	return err
}

func (u *userService) GenerateToken(userUuid string) (td *TokenDetail, err error) {
	td = &TokenDetail{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_uuid"] = userUuid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(u.atSecret))
	if err != nil {
		return nil, err

	}
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_uuid"] = userUuid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(u.rtSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (u *userService) CreateAuth(uuid string, td *TokenDetail) (err error) {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	err = u.repo.AddTokenToClient(td.AccessToken, uuid, at.Sub(now))
	if err != nil {
		return err
	}
	err = u.repo.AddTokenToClient(td.RefreshToken, uuid, rt.Sub(now))
	if err != nil {
		return err
	}
	return nil
}
