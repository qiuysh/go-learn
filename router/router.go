package router

import (
	"go-learn/controllers"
	"go-learn/middlewares"
	"github.com/gin-gonic/gin"
)

func RouterConfig(r *gin.Engine) {

	GroupV1 := r.Group("/gin/v1")
	
	{
		GroupV1.POST("/login", controllers.Login)
	}

	GroupV2 := r.Group("/gin/v1")

	GroupV2.Use(middlewares.Authorized())

	{
		GroupV2.POST("/getUserList", controllers.Users)
		GroupV2.POST("/getRoleList", controllers.Roles)
		GroupV2.POST("/getMenuList", controllers.Menus)
	}
}
