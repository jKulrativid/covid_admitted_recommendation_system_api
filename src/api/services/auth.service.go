package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	ExtractMetadata(jwtToken string) (*AccessDetails, error)
	FetchAuth(authD *AccessDetails) (string, error)
}

type authService struct {
	repo      repositories.AuthRepo
	jwtSecret string
}

type TokenDetail struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	TokenUuid string
	UserUuid  string
}

func NewAuthService(r repositories.AuthRepo) AuthService {
	sc := os.Getenv("JWT_SECRET")
	if sc == "" {
		log.Fatal("no environment Variable \"JWT_SECRET\" given")
	}
	return &authService{
		repo:      r,
		jwtSecret: sc,
	}
}

func (a *authService) verifyToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, entities.ErrorInvalildToken
		}
		return []byte(a.jwtSecret), nil
	})
	if err != nil {
		return nil, entities.ErrorExpiredToken
	}
	return token, nil
}

func (a *authService) ExtractMetadata(jwtToken string) (*AccessDetails, error) {
	token, err := a.verifyToken(jwtToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		tokenUuid, ok := claims["token_uuid"].(string)
		if !ok {
			return nil, entities.ErrorInvalildToken
		}
		uuid, ok := claims["user_uuid"].(string)
		if !ok {
			return nil, entities.ErrorInvalildToken
		}
		return &AccessDetails{
			TokenUuid: tokenUuid,
			UserUuid:  uuid,
		}, nil
	}
	return nil, entities.ErrorInvalildToken
}

func (a *authService) FetchAuth(authD *AccessDetails) (string, error) {
	uid, err := a.repo.GetFromClient(authD.TokenUuid)
	if err != nil {
		return "", err
	}
	return uid, nil
}
