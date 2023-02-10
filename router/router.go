package router

import (
	"react-system/controller"

	"github.com/gin-gonic/gin"
)

func RouterConfig(r *gin.Engine) {
	GroupV1 := r.Group("/rs/v1")
	{
		GroupV1.Any("/login", controller.Login)
		GroupV1.Any("/getUserList", controller.Users)
	}
}
