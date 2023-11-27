package config

import (
	"fmt"
	"gorm.io/gorm"
  "github.com/spf13/viper"

)


const (
	PORT      = ":8081"
	JWTSECRET = "123456"
)

var (
	GDB 	*gorm.DB
	CONF  *ConfigYAMLStruct
)


type DataBaseStruct struct {
	Username 		string 		`yaml: "username"`
	Password 		string 		  `yaml: "password"`
	Host 				string 		`yaml: "host"`
	Port 				string 		  `yaml: "port"`
	Databasename 			string 		`yaml: "databasename"`
}

type ConfigYAMLStruct struct {
	Database 		  DataBaseStruct 		`yaml: "database"  json: "database"`
}

func init()  {
	
	viper.SetConfigFile("./conf.yaml")
	// 寻找配置文件并读取
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %w", err))
	}

	// 将读取的配置信息保存至全局变量 Conf
	if err := viper.Unmarshal(&CONF); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
}