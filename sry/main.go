package main

import (
	"context"

	"sry/setting"
	"sry/router"
	"sry/logger"
	"sry/mysql"
	"sry/redis"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {	
	//加载配置 setings 参见viper小节
	setting.SettingInit()
	//初始化日志
	logger.InitLogger() 
	//初始化mysql连接
	db := mysql.InitMysql()
	defer db.Close()
	//初始化redis连接
	ctx := context.Background()
	rdb, err := redis.InitRedisClient(ctx)
	defer rdb.Close()
	if err != nil {
		panic(err)
	}
	//注册路由
	r := router.Setup()
	r.Run()
	//启动服务 优雅关机

	//初始化之后记得 defer or close
}
