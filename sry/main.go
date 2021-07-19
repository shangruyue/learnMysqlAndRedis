package main

import (
	//"context"

	"fmt"
	"sry/logger"
	"sry/router"
	"sry/setting"

	"sry/dao/mysql"
	//"sry/redis"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {	
	//加载配置 setings 参见viper小节
	setting.SettingInit()
	//初始化日志
	logger.SetupSugarLogger() 
	//初始化mysql连接
	db := mysql.InitMysql()
	defer db.Close()
	//初始化redis连接
	// ctx := context.Background()
	// err := redis.InitRedisClient(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// redis.MyredisClose()
	//注册路由
	r := router.Setup()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("run server failed: %v", err)
	}
	//启动服务 优雅关机

	//初始化之后记得 defer or close
}
