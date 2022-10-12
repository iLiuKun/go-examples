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

	//设置一个name对应的值是hello
	err = redisDb.Set("name", "hello", 0).Err()
	if err != nil {
		panic(err)
	}
	redisDb.Append("name", "China")
	val,err:=redisDb.Get("name").Result()
	fmt.Println(val) //helloChina
	
	//设置key的过期时间为5秒
	redisDb.Expire("name", 5000000000)
	//批量删除
	val2,err2:=redisDb.Del("name1", "name2", "key1", "key2").Result()
	if err != nil {
		panic(err2)
	}
	fmt.Println(val2)
}

