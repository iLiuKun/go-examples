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
	//设置一个age测试自增、自减
	err = redisDb.Set("age", "20", 0).Err()
	if err != nil {
		panic(err)
	}
	redisDb.Incr("age")      // 自增
	redisDb.IncrBy("age", 5) //+5
	redisDb.Decr("age")      // 自减
	redisDb.DecrBy("age", 3) //-3 此时age的值是22

	var val string
	val, err= redisDb.Get("age").Result()
	fmt.Println("age=",val) //22
}
