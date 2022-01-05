package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Validate(u interface{}) error
	Register(newUser *entities.UserRegister) error
	SignIn(user *entities.UserSignIn) (string, error)
	SignOut(user *entities.UserSignIn) (string, error)
	GenerateToken(userUuid string) (*TokenDetail, error)
	CreateAuth(uid string, td *TokenDetail) error
	DeleteAuth(uid string) error
	userExist(userName, email string) bool
}

type userService struct {
	repo      repositories.UserRepository
	validator *validator.Validate
	jwtSecret string
}

func NewUserService(r repositories.UserRepository) UserService {
	sc := os.Getenv("JWT_SECRET")
	if sc == "" {
		log.Fatal("no environment Variable \"JWT_SECRET\" given")
	}
	return &userService{
		repo:      r,
		validator: validator.New(),
		jwtSecret: sc,
	}
}

func (u *userService) userExist(userName, email string) bool {
	if err := u.repo.GetUserFromEmail(&entities.User{}, email); err == nil {
		return true
	}
	if err := u.repo.GetUserFromUserName(&entities.User{}, userName); err == nil {
		return true
	}
	return false
}

func (u *userService) Register(newUser *entities.UserRegister) error {
	if u.userExist(newUser.UserName, newUser.Email) {
		return entities.ErrorConflict
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	user := &entities.User{
		Uid:            uuid.NewV4().String(),
		UserName:       newUser.UserName,
		Email:          newUser.Email,
		HashedPassword: string(hashedPassword),
	}
	if err := u.repo.CreateNewUser(user); err != nil {
		return err
	}
	if err := u.repo.CreateNewUserDirectory(user.Uid); err != nil {
		return err
	}
	return nil
}

func (u *userService) SignIn(userSignIn *entities.UserSignIn) (string, error) {
	var user entities.User
	if err := u.repo.GetUserFromUserName(&user, userSignIn.UserName); err != nil {
		return "", entities.ErrorUnAuthorized
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userSignIn.Password)); err != nil {
		return "", entities.ErrorUnAuthorized
	}
	return user.Uid, nil
}

func (u *userService) SignOut(userSignIn *entities.UserSignIn) (string, error) {
	var user entities.User
	if err := u.repo.GetUserFromUserName(&user, userSignIn.UserName); err != nil {
		return "", entities.ErrorUnAuthorized
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userSignIn.Password)); err != nil {
		return "", entities.ErrorUnAuthorized
	}
	return user.Uid, nil
}

func (u *userService) GenerateToken(userUuid string) (*TokenDetail, error) {
	td := &TokenDetail{}
	var err error

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["token_uuid"] = td.AccessUuid
	atClaims["user_uuid"] = userUuid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return nil, err
	}
	rtClaims := jwt.MapClaims{}
	rtClaims["token_uuid"] = td.RefreshUuid
	rtClaims["user_uuid"] = userUuid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (u *userService) CreateAuth(uid string, td *TokenDetail) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	err := u.repo.SaveJWT(td.AccessUuid, uid, at.Sub(now))
	if err != nil {
		return err
	}
	err = u.repo.SaveJWT(td.RefreshUuid, uid, rt.Sub(now))
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) DeleteAuth(uid string) error {
	return u.repo.DeleteJWT(uid)
}

func (u *userService) Validate(i interface{}) error {
	if err := u.validator.Struct(i); err != nil {
		return entities.ErrorInvalidUserForm
	}
	return nil
}
