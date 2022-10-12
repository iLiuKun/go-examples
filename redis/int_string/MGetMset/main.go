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

	//批量设置key1对应的值为value1，key2对应的值为value2，key3对应的值为value3，
	err = redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3").Err()
	if err != nil {
		panic(err)
	}

	// MGet函数可以传入任意个key，一次性返回多个值。
	// 这里Result返回两个值，第一个值是一个数组，第二个值是错误信息
	vals, err := redisDb.MGet("key1", "key2", "key3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals) //[value1 value2 value3]
}

