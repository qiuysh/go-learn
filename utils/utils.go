package utils

import (
	"fmt"
	"time"
	"go-learn/config"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

func GetFromPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密失败")
		return ""
	}
	return string(hash)
}


func ComparePwd(hash string, pwd string) bool {
 err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
 if err == nil {
   return true
 }
 return false
}
