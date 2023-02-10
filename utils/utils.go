package utils

import (
	"fmt"
	"time"

	"react-system/config"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id       int16  `json:"id"`
	Username string `json:"username"`
}

type StandardClaims struct {
	*User
	jwt.StandardClaims
}

var jwtSecret = []byte(config.JWTSECRET)

func GenerateToken(id int16, username string) (string, error) {

	expire := time.Now().Add(time.Hour).Unix()
	user := User{
		id,
		username,
	}
	claims := StandardClaims{
		&user,
		jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "admin",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(tokenString string) (User, error) {
	claimsUser := &StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claimsUser, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	return *claimsUser.User, err
}
