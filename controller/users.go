package controller

import (
	"net/http"
	"react-system/config"
	"react-system/enity"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Id       int16  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func Users(c *gin.Context) {
	var userList []UserInfo
	// post
	res := enity.Result{}
	rows, err := config.DB.Query("select id, username, email, phone from user limit ?, ?", 0, 9)
	if err != nil {
		res.SetSuccess(enity.RESULT_FAILURE)
		res.SetMessage("账号或密码错误！")
		c.JSON(http.StatusOK, res)
	} else {
		for rows.Next() {
			var userInfo UserInfo
			err = rows.Scan(&userInfo.Id, &userInfo.Username, &userInfo.Email, &userInfo.Phone)
			if err != nil {
				return
			}
			userList = append(userList, userInfo)
		}
		res.SetSuccess(enity.RESULT_SUCCESS)
		res.SetMessage("请求成功")
		res.SetData(userList)
		c.JSON(http.StatusOK, res)
	}
	c.Abort()
}
