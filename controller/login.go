package controller

import (
	"fmt"
	"go-learn/config"
	"go-learn/enity"
	"go-learn/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int16  `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user User
	// post
	res := enity.Result{}
	if err := c.ShouldBind(&user); err == nil {
		err := config.DB.QueryRow("select id, username, password from user where username = ? and password = ?", user.Username, user.Password).Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			res.SetSuccess(enity.RESULT_FAILURE)
			res.SetMessage("账号或密码错误！")
			c.JSON(http.StatusOK, res)
		} else {
			token, err := utils.GenerateToken(user.Id, user.Username)
			if err != nil {
				fmt.Println("生成token失败")
			}
			result := map[string]interface{}{
				"id":       user.Id,
				"username": user.Username,
				"token":    token,
			}
			res.SetSuccess(enity.RESULT_SUCCESS)
			res.SetMessage("请求成功")
			res.SetData(result)
			c.JSON(http.StatusOK, res)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Abort()
}
