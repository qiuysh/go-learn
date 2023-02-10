package main

import (
	"fmt"
	"go-learn/config"
	"go-learn/router"

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
