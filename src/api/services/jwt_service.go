package services

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	secret       string
	issuer       string
	expireSecond int64
}

func NewJWTService(issuer string, expireSecond int64) *JWTService {
	secret := os.Getenv("SECRET_JWT")
	if secret == "" {
		log.Fatal("Crashed in NewJWTService (package services) : No Environment Variable \"SECRET_JWT\" Given")
	}
	return &JWTService{
		secret:       secret,
		issuer:       issuer,
		expireSecond: expireSecond,
	}
}

func (service *JWTService) GenerateToken(userID string) string {
	atClaims := jwt.MapClaims{}
	atClaims["userID"] = userID
	return ""

}

func (service *JWTService) VerifyToken(jwtToken string) bool {
	return false

}
