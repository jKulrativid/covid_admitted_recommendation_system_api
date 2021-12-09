package services

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	secret string
	claims jwt.StandardClaims
}

func NewJWTService() *JWTService {
	var jwtService JWTService
	secret := os.Getenv("SECRET_JWT")
	if secret == "" {
		log.Fatal("Crashed : No Environment Variable \"SECRET_JWT\" Given")
	}
}
