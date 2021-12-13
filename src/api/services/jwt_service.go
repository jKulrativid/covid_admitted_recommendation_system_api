package services

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type JWTService struct {
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

func NewJWTService() *JWTService {
	as := os.Getenv("ACCESS_SECRET_JWT")
	rs := os.Getenv("REFRESH_SECRET_JWT")
	if as == "" || rs == "" {
		log.Fatal("Crashed in NewJWTService (jwt_service.go) : No Environment Variable \"SECRET_JWT\" Given")
	}
	return &JWTService{
		atSecret: as,
		rtSecret: rs,
	}
}

func (service *JWTService) GenerateToken(userID uint64) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(service.atSecret))
	if err != nil {
		return nil, err

	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(service.rtSecret))
	if err != nil {
		return nil, err

	}

	return td, nil

}

func (service *JWTService) VerifyToken(jwtToken string) error {
	return nil

}
