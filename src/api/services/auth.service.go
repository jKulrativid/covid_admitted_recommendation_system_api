package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	VerifyToken(jwtToken string) (*jwt.Token, error)
	TokenValid(jwtToken string) error
	ExtractMetadata(jwtToken string) (*AccessDetails, error)
	FetchAuth(authD *AccessDetails) (string, error)
}

type authService struct {
	repo     repositories.AuthRepo
	atSecret string
	rtSecret string
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
	AccessUuid string
	UserUuid   string
}

func NewAuthService(r repositories.AuthRepo) AuthService {
	as := os.Getenv("ACCESS_JWT_SECRET")
	rs := os.Getenv("REFRESH_JWT_SECRET")
	if as == "" || rs == "" {
		log.Fatal("Crashed in NewJWTService (jwt_service.go) : No Environment Variable \"ACCESS_JWT_SECRET\" or \"REFRESH_JWT_SECRET\" Given")
	}
	return &authService{
		repo:     r,
		atSecret: as,
		rtSecret: rs,
	}
}

func (a *authService) VerifyToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, entities.ErrorInvalildToken
		}
		return []byte(a.atSecret), nil
	})
	if err != nil {
		return nil, entities.ErrorExpiredToken
	}
	return token, nil

}

func (a *authService) TokenValid(jwtToken string) error {
	token, err := a.VerifyToken(jwtToken)
	if err != nil {
		return err
	}
	if !token.Valid {
		return entities.ErrorInvalildToken
	}
	return nil
}

func (a *authService) ExtractMetadata(jwtToken string) (*AccessDetails, error) {
	token, err := a.VerifyToken(jwtToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, entities.ErrorInvalildToken
		}
		uuid, ok := claims["user_uuid"].(string)
		if !ok {
			return nil, entities.ErrorInvalildToken
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserUuid:   uuid,
		}, nil
	}
	return nil, entities.ErrorInvalildToken
}

func (a *authService) FetchAuth(authD *AccessDetails) (string, error) {
	uid, err := a.repo.GetFromClient(authD.AccessUuid)
	if err != nil {
		return "", err
	}
	return uid, nil
}
