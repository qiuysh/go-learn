package controllers

import (
	global "go-learn/config"
	"go-learn/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Id       int16  `json:"id" gorm:"column:id; primaryKey"`
	Username string `json:"username" gorm:"column:username"`
	Updated_at string `json:"updated_at" gorm:"column:updated_at"`
}

func Users(c *gin.Context) {
	var userList = make([]UserInfo, 0)
	// post
	res := utils.Result{}
	
	rows, err := global.GDB.Table("goadmin_users").Limit(10).Rows()

	defer rows.Close()

	if err != nil {
		res.SetSuccess(utils.RESULT_FAILURE)
		res.SetMessage("请求数据错误！")
		c.JSON(http.StatusOK, res)
	} else {
		for rows.Next() {
			var userInfo UserInfo
			err := global.GDB.ScanRows(rows, &userInfo)
			if err != nil {
				return
			}
			userList = append(userList, userInfo)
		}
		res.SetSuccess(utils.RESULT_SUCCESS)
		res.SetMessage("请求成功")
		res.SetData(userList)
		c.JSON(http.StatusOK, res)
	}
	c.Abort()
}
