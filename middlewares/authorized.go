package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-learn/utils"
) 

func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {

		var authToken = c.Request.Header.Get("token")

		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
				"data":   nil,
			})
			c.Abort()
			return
		}

		_, err := utils.ParseToken(authToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": -1,
				"msg":    "请求携带的token已失效",
				"data":   nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
