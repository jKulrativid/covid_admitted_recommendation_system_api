package services

import (
	"covid_admission_api/repositories"
	"fmt"
	"log"
	"os"
	"strings"

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

func (a *authService) ExtractToken(jwtToken string) string {
	strArr := strings.Split(jwtToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (a *authService) VerifyToken(jwtToken string) (*jwt.Token, error) {
	tokenString := a.ExtractToken(jwtToken)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.atSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (a *authService) TokenValid(jwtToken string) error {
	token, err := a.VerifyToken(jwtToken)
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
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
			return nil, fmt.Errorf("token payload invalid")
		}
		uuid, ok := claims["user_uuid"].(string)
		if !ok {
			return nil, fmt.Errorf("token payload invalid")
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserUuid:   uuid,
		}, nil
	}
	return nil, err
}

func (a *authService) FetchAuth(authD *AccessDetails) (string, error) {
	uuid, err := a.repo.GetFromClient(authD.AccessUuid)
	if err != nil {
		return "", err
	}
	return uuid, nil
}
