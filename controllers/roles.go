package controllers

import (
	global "go-learn/config"
	"go-learn/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

type RoleInfo struct {
	Id       int16  `json:"id" gorm:"column:id; primaryKey"`
	Name string `json:"name" gorm:"column:name"`
	Updated_at string `json:"updated_at" gorm:"column:updated_at"`
}

func Roles(c *gin.Context) {
	var roleList = make([]RoleInfo, 0)
	// post
	res := utils.Result{}
	
	rows, err := global.GDB.Table("goadmin_roles").Limit(10).Rows()

	defer rows.Close()

	if err != nil {
		res.SetSuccess(utils.RESULT_FAILURE)
		res.SetMessage("请求数据错误！")
		c.JSON(http.StatusOK, res)
	} else {
		for rows.Next() {
			var roleInfo RoleInfo
			err := global.GDB.ScanRows(rows, &roleInfo)
			if err != nil {
				return
			}
			roleList = append(roleList, roleInfo)
		}
		res.SetSuccess(utils.RESULT_SUCCESS)
		res.SetMessage("请求成功")
		res.SetData(roleList)
		c.JSON(http.StatusOK, res)
	}
	c.Abort()
}
