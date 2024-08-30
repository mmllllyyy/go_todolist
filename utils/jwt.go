package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("todolist")

type Claims struct {
	UserID   uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(id uint, username string) (string, error) {
	c := Claims{
		UserID:   id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    "todolist",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
