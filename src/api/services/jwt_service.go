package services

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

type AccessDetails struct {
	AccessUuid string
	UserUuid   string
}

func NewJWTService() *JWTService {
	as := os.Getenv("ACCESS_JWT_SECRET")
	rs := os.Getenv("REFRESH_JWT_SECRET")
	if as == "" || rs == "" {
		log.Fatal("Crashed in NewJWTService (jwt_service.go) : No Environment Variable \"ACCESS_JWT_SECRET\" or \"REFRESH_JWT_SECRET\" Given")
	}
	return &JWTService{
		atSecret: as,
		rtSecret: rs,
	}
}

func (service *JWTService) GenerateToken(userUuid string) (td *TokenDetail, err error) {
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
	td.AccessToken, err = at.SignedString([]byte(service.atSecret))
	if err != nil {
		return nil, err

	}
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_uuid"] = userUuid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(service.rtSecret))
	if err != nil {
		return nil, err

	}
	return td, nil

}

func (service *JWTService) extractToken(jwtToken string) string {
	strArr := strings.Split(jwtToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (service *JWTService) VerifyToken(jwtToken string) (*jwt.Token, error) {
	tokenString := service.extractToken(jwtToken)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (service *JWTService) TokenValid(jwtToken string) error {
	token, err := service.VerifyToken(jwtToken)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (service *JWTService) ExtractMetadata(jwtToken string) (*AccessDetails, error) {
	token, err := service.VerifyToken(jwtToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		uuid := claims["user_uuid"].(string)
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserUuid:   uuid,
		}, nil
	}
	return nil, err
}
