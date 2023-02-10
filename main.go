package main

import (
	"fmt"
	"react-system/config"
	"react-system/router"
	"react-system/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	router.RouterConfig(engine)
	config.InitDB()
	err := engine.Run(config.PORT)
	if err != nil {
		fmt.Println(err.Error())
	}

}
