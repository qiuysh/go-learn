package controllers

import (
	"fmt"
	global "go-learn/config"
	"go-learn/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)


type User struct {
	Id       int16  `json:"id" gorm:"column:id; primaryKey"`
	Username string `json:"username"  binding:"required"  gorm:"column:username"`
	Password string `json:"password" binding:"required"  gorm:"column:password"`
}

func Login(c *gin.Context) {
	var user User
	// post
	res := utils.Result{}

	if err := c.ShouldBindJSON(&user); err == nil {

		var pwdtext string = user.Password

		global.GDB.Table("goadmin_users").Where("username = ?", user.Username).Take(&user)

		if utils.ComparePwd(user.Password, pwdtext) {

			token, err := utils.GenerateToken(user.Id, user.Username)
			if err != nil {
				fmt.Println("生成token失败")
			}
			result := map[string]interface{}{
				"id":       user.Id,
				"username": user.Username,
				"token":    token,
			}
			res.SetSuccess(utils.RESULT_SUCCESS)
			res.SetMessage("请求成功")
			res.SetData(result)
			c.JSON(http.StatusOK, res)

		} else {

			res.SetSuccess(utils.RESULT_FAILURE)
			res.SetMessage("账号或密码错误！")
			c.JSON(http.StatusOK, res)
			
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	c.Abort()
}
