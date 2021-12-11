package services

import (
	"log"
	"os"
)

type JWTService struct {
	secret       string
	issuer       string
	expireSecond int64
}

func NewJWTService(issuer string, expireSecond int64) *JWTService {
	secret := os.Getenv("SECRET_JWT")
	if secret == "" {
		log.Fatal("Crashed in NewJWTService (jwt_service.go) : No Environment Variable \"SECRET_JWT\" Given")
	}
	return &JWTService{
		secret:       secret,
		issuer:       issuer,
		expireSecond: expireSecond,
	}
}

func (service *JWTService) GenerateToken(userID string) string {
	return ""

}

func (service *JWTService) VerifyToken(jwtToken string) error {
	return nil

}
