package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的redisDb变量
var redisDb *redis.Client

// 根据redis配置初始化一个客户端
func initClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initClient()
	if err != nil {
		//redis连接错误
		panic(err)
	}
	// 第三个参数代表key的过期时间，0代表不会过期。
	err = redisDb.Set("name1", "zhangsan", 0).Err()
	if err != nil {
		panic(err)
	}
	var val string
	// Result函数返回两个值，第一个是key的值，第二个是错误信息
	val, err = redisDb.Get("name1").Result()
	// 判断查询是否出错
	if err != nil {
		panic(err)
	}
	fmt.Println("name1的值：", val) //name1的值：zhangsan
}

