package middleware

import (
	"fmt"
	"time"

	"react-system/config"

	"github.com/dgrijalva/jwt-go"
)


func Authorized() gin.HandlerFunc {
	return func (c *gin.Context)  {
		var authToken = c.Responese.Header("")
	}
}

