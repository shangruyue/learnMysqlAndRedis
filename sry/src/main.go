package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"sry/src/setting"

	gin "github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
)

var db *sql.DB

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient(ctx context.Context) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root", // set password is "root"
		DB:       0,  // use default DB
		PoolSize: 100, //连接池大小
	})

	//这一行是为了测试git做的修改

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Printf("redis connect success...\n")
	return nil
}

//redis set get example
func redisSetAndGetExample(){
	ctx := context.Background()
	if err := initClient(ctx); err != nil {
		panic(err)
	}
	rdb.Set(ctx, "xiaoman", 22, time.Second*5)

	result := rdb.Get(ctx, "xiaoman")
	result1 := rdb.Get(ctx, "xiaoman1")
	fmt.Printf(result.String())
	fmt.Println()
	fmt.Printf(result1.String())
	fmt.Println()
}

func main() {
	setting.Init()
	//加载配置 setings 参见viper小节
	//初始化日志 
	//初始化mysql连接
	//初始化redis连接
	//注册路由
	//启动服务 优雅关机

	//初始化之后记得 defer or close

	r := gin.Default() //创建一个默认的路由引擎
	r.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "get Book",
		})
	})
	r.POST("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "post Book",
		})
	})
	r.PUT("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "put Book",
		})
	})
	r.DELETE("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "delete Book",
		})
	})
	r.Run()
	// ctx := context.Background()
	// err := initClient(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	//redisSetAndGetExample()
	//mysqlStr := "postgres://postgres:AWc7nCRAeesaa3d7C6NS@uservicedb-dev.cyccjyevtjgf.us-west-2.rds.amazonaws.com:5432/uservicedbdev"
	//mysqlStr := "root:root@tcp(127.0.0.1:3306)/mysql"
	////db, err := sql.Open("pgx", mysqlStr)
	//db, err := sql.Open("mysql", mysqlStr)
	//if err != nil {
	//	fmt.Printf("Open db failed: %v\n", err)
	//}
	//err = db.Ping()
	//if err != nil {
	//	fmt.Printf("Ping failed: %v\n", err)
	//}
	//defer db.Close()
	//fmt.Printf("connect db success...")
}
