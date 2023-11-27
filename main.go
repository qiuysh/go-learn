package main

import (
	"fmt"
	"go-learn/utils"
	"go-learn/config"
	"go-learn/router"
	"github.com/gin-gonic/gin"
)

func main() {

	utils.InitDBConfig()
	
	gin.SetMode(gin.DebugMode)

	engine := gin.Default()

	router.RouterConfig(engine)

	err := engine.Run(config.PORT)

	if err != nil {
		fmt.Println(err.Error())
	}
}
