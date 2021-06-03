package redis

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func InitRedisClient(ctx context.Context) (err error) {
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
	if err := InitRedisClient(ctx); err != nil {
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

func MyredisClose() {
	defer rdb.Close()
}