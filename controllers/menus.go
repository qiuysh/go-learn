package controllers

import (
	global "go-learn/config"
	"go-learn/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

type MenuInfo struct {
	Id       		  int16  		`json:"id" gorm:"column:id; primaryKey"`
	Parent_id     int16  		`json:"parent_id" gorm:"column:parent_id"`
	Title 				string 		`json:"title" gorm:"column:title"`
	Icon 					string 		`json:"icon" gorm:"column:icon"`
	Uri 					string 		`json:"uri" gorm:"column:uri"`
	Updated_at 		string 		`json:"updated_at" gorm:"column:updated_at"`
}

func Menus(c *gin.Context) {
	var menuList = make([]MenuInfo, 0)
	// post
	res := utils.Result{}
	
	rows, err := global.GDB.Table("goadmin_menu").Limit(10).Rows()

	defer rows.Close()

	if err != nil {
		res.SetSuccess(utils.RESULT_FAILURE)
		res.SetMessage("请求数据错误！")
		c.JSON(http.StatusOK, res)
	} else {
		for rows.Next() {
			var MenuInfo MenuInfo
			err := global.GDB.ScanRows(rows, &MenuInfo)
			if err != nil {
				return
			}
			menuList = append(menuList, MenuInfo)
		}
		res.SetSuccess(utils.RESULT_SUCCESS)
		res.SetMessage("请求成功")
		res.SetData(menuList)
		c.JSON(http.StatusOK, res)
	}
	c.Abort()
}
