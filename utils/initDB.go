package utils

import (
	"fmt"
	"time"
	"strings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	global "go-learn/config"
)


// 注意方法名大写，就是public
func InitDBConfig() {

	var _config global.DataBaseStruct  = global.CONF.Database

	
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	dsn := strings.Join([]string{_config.Username, ":", _config.Password, "@tcp(", _config.Host, ":", _config.Port, ")/", _config.Databasename, "?charset=utf8"}, "")

	fmt.Println(dsn)

	//打开数据库
	global.GDB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	sqlDB, _ := global.GDB.DB()

	//验证连接
	if err := sqlDB.Ping(); err != nil {
		fmt.Println("数据库连接失败!")
		return
	}

	fmt.Println("数据库连接成功!")

	//设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)   
	
	//设置数据库连接池参数
	sqlDB.SetMaxIdleConns(20) 

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//延时关闭数据库连接
	// defer sqlDB.Close()

}